#!/usr/bin/env python3
"""Audita o catálogo do protoncommand contra as fontes primárias do upstream.

Dois modos, independentes:

  --offline   invariantes locais (rápido, sem rede). É o que o CI roda.
  --online    extrai toda env var do catálogo e cruza com o script do Proton,
              os READMEs e CHANGELOGs do Proton-CachyOS e do GE, e com o
              README/config do DXVK. Informa o que não encontrou, não falha.

Um env var sem menção upstream não é erro: DXVK_*, MANGOHUD*, VKD3D_*,
LOW_LATENCY_LAYER* e friends pertencem a outros projetos. Por isso o relatório
agrupa por prefixo e diz onde cada um deveria ser conferido. O que o script
acha mesmo é renomeado, obsoleto ou simplesmente inventado.

    tools/auditar-catalogo.py --offline
    tools/auditar-catalogo.py --online
"""

from __future__ import annotations

import argparse
import collections
import os
import re
import subprocess
import sys

RAIZ = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
CATALOGO = os.path.join(RAIZ, "commands.go")
CACHE = os.environ.get("AUDIT_CACHE") or os.path.join(RAIZ, ".auditoria-cache")

# prefixo -> onde a variável deveria ser conferida
DESTINO = [
    ("PROTON_", "script do Proton (ValveSoftware/Proton) + README/CHANGELOG do Proton-CachyOS e GE"),
    ("WINE_", "script do Proton e/ou patches do Wine"),
    ("DXVK_", "README e dxvk.conf do doitsujin/dxvk, e os forks (low-latency, nvapi, sarek)"),
    ("VKD3D_", "README do HansKristian-Work/vkd3d-proton e do netborg-afps/vkd3d-low-latency"),
    ("MANGOHUD", "README do MangoHud"),
    ("GAMEMODERUN", "README do GameMode"),
    ("GAMESCOPE", "README do gamescope"),
    ("LOW_LATENCY_LAYER", "README do Korthos-Software/low_latency_layer"),
    ("LSFGVK", "docs/Configuration.md do PancakeTAS/lsfg-vk"),
    ("DXIL_", "DXIL-SPIR-V do vkd3d-proton"),
    ("FSR_", "patches de FSR do Proton"),
    ("MESA_", "fontes do Mesa (src/vulkan/wsi, src/amd/vulkan)"),
    ("ENABLE_LAYER_MESA_", "VkLayer_MESA_anti_lag.json.in no Mesa"),
    ("DISABLE_LAYER_MESA_", "VkLayer_MESA_anti_lag.json.in no Mesa"),
    ("RADV_", "radv_instance.c (radv_debug_options) no Mesa"),
    ("ENABLE_HDR_WSI", "camada de HDR do driver Vulkan"),
    ("DRI_PRIME", "PRIME do Mesa, não é variável do Proton"),
    ("HOST_LC_ALL", "script do Proton e locale do Wine"),
    ("WINEDLLOVERRIDES", "Wine"),
]

FONTES = {
    "proton_11.0": "https://raw.githubusercontent.com/ValveSoftware/Proton/proton_11.0/proton",
    "cachyos_readme": "https://raw.githubusercontent.com/CachyOS/proton-cachyos/master/README.md",
    "cachyos_changelog": "https://raw.githubusercontent.com/CachyOS/proton-cachyos/master/CHANGELOG.md",
    "ge_readme": "https://raw.githubusercontent.com/GloriousEggroll/proton-ge-custom/master/README.md",
    "dxvk_readme": "https://raw.githubusercontent.com/doitsujin/dxvk/master/README.md",
    "dxvk_conf": "https://raw.githubusercontent.com/doitsujin/dxvk/master/dxvk.conf",
    "low_latency_layer": "https://raw.githubusercontent.com/Korthos-Software/low_latency_layer/main/README.md",
    "dxvk_low_latency": "https://raw.githubusercontent.com/netborg-afps/dxvk-low-latency/master/README.md",
    "vkd3d_low_latency": "https://raw.githubusercontent.com/netborg-afps/vkd3d-low-latency/master/README.md",
    "lsfgvk": "https://raw.githubusercontent.com/xXJSONDeruloXx/lsfg-vk-arm64/v2-arm64-port/docs/Configuration.md",
}

# variável -> o que o upstream diz que aconteceu com ela
REMOVIDAS = {
    "DXVK_FRAME_RATE": "removida no DXVK 3.0 (release notes); use dxvk.maxFrameRate via DXVK_CONFIG",
    "VKD3D_FRAME_RATE": "não existe no vkd3d-proton upstream, só no fork vkd3d-low-latency",
    "PROTON_USE_D9VK": "obsoleta no Proton 5.0",
    "PROTON_USE_SECCOMP": "obsoleta no Proton 5.13",
    "PROTON_MEDIA_USE_GST": "renomeada para PROTON_MEDIA_FORCE_GST",
    "PROTON_FSR3_UPGRADE": "renomeada para PROTON_FFX3_UPGRADE no CachyOS 11+",
    "PROTON_FSR4_RDNA3_UPGRADE": "removida no CachyOS 11 (wmma_rdna3_workaround deixou de ser necessário)",
}

PADRAO_VAR = re.compile(r"\b([A-Z][A-Z0-9_]{2,})\b")
IGNORAR = {
    "PT", "EN", "APPIMAGE", "AMD", "CPU", "DLC", "DPI", "FPS", "GPU", "HDR", "HLSL",
    "INT", "RTX", "SDK", "UE", "VKD3D", "VULKAN", "VULKANAPI", "WINE", "WINEPREFIX",
    "X86", "XBOX", "YUV", "CPU", "ACP", "AO", "API", "D3D", "D3D9", "D3D10", "D3D11",
    "D3D12", "DXGI", "DX9", "DX10", "DX11", "DX12", "FFX3", "FFX4", "CS", "LAA",
}


def ler(caminho: str) -> str:
    with open(caminho, encoding="utf-8", errors="replace") as f:
        return f.read()


def baixar(nome: str, url: str, cache: str) -> str:
    os.makedirs(cache, exist_ok=True)
    destino = os.path.join(cache, nome)
    if os.path.exists(destino) and os.path.getsize(destino) > 0:
        return ler(destino)
    try:
        out = subprocess.run(
            ["curl", "-sfL", "--max-time", "40", url],
            capture_output=True, text=True, timeout=60, check=True,
        ).stdout
    except (subprocess.SubprocessError, OSError) as e:
        print(f"aviso: não consegui baixar {nome} ({e.__class__.__name__}); seguindo sem ele",
              file=sys.stderr)
        return ""
    with open(destino, "w", encoding="utf-8") as f:
        f.write(out)
    return out


STRING_GO = re.compile(r'Command:\s+"((?:[^"\\]|\\.)*)"')


def desescapar(valor: str) -> str:
    """Converte o literal de string do Go para o valor efetivo."""
    return (valor.replace('\\\\', "\x00")
                 .replace('\\"', '"')
                 .replace("\\\\", "\\")
                 .replace("\x00", "\\"))


def entradas_do_catalogo() -> list[dict]:
    """Extrai Command/Title/Category/Compat/Description do commands.go."""
    fonte = ler(CATALOGO)
    blocos = re.split(r"\n\t\t\{\n", fonte)
    achados = []
    for b in blocos[1:]:
        cmd = re.search(STRING_GO, b)
        if not cmd:
            continue
        def campo(nome: str) -> tuple[str, str]:
            m = re.search(nome + r': Localized\{\s*PT: "(.*?)",\s*EN: "(.*?)",', b, re.S)
            return (m.group(1), m.group(2)) if m else ("", "")
        achados.append({
            "command": desescapar(cmd.group(1)),
            "title": campo("Title"),
            "category": campo("Category"),
            "compat": campo("Compat"),
            "description": campo("Description"),
        })
    return achados


def env_vars(comando: str) -> list[str]:
    """Tokens de assignment no início dos tokens, ignorando %command% e flags."""
    found = []
    for tok in comando.replace("%command%", " ").split():
        if "=" in tok and not tok.startswith("="):
            k = tok.split("=", 1)[0]
            if re.fullmatch(r"[A-Z][A-Z0-9_]+", k):
                found.append(k)
    return found


def destino_de(var: str) -> str:
    for prefixo, onde in DESTINO:
        if var.startswith(prefixo):
            return onde
    return "fonte não mapeada: adicione em DESTINO no auditor"


# --------------------------------------------------------------------------- offline

# Resto de edição em texto: aspas órfãos e asterisco colado em letra.
# Já apareceu "o jogo se*''confundir" em commands.go por um replace mal
# feito, e nenhum teste pegou — o texto é lido por gente, não compilado.
GLITCH = re.compile(r"[A-Za-z]''|\*''|[A-Za-z]\*[A-Za-z']")

# Inflexão gramatical inglesa dentro do texto em português. Já apareceu
# "as duas continuam controlling as mesmas versões". Termo técnico de
# área (timing, pacing, spoofing, clipping, rendering) é aceito: a lista
# é só de verbo em inglês conjugado, que em PT sairia como gerúndio.
INGLES_NO_PT = re.compile(
    r"\b(controlling|enabling|disabling|running|using|setting|getting|"
    r"turning|showing|making|giving|taking|leaving|allowing|forcing|"
    r"replacing|removing|adding|meaning|working|looking|keeping|"
    r"putting|calling|reading|writing)\b",
    re.I,
)

# Palavra gramatical portuguesa dentro do texto em inglês, o caminho
# inverso do erro anterior. Identificador técnico não entra aqui.
# O sufixo -ção/-ções precisa de \w* antes: \b não existe à esquerda do
# ç, então a forma simples casa só a palavra isolada e nunca aparece.
PT_NO_EN = re.compile(
    r"\b\w*(?:ção|ções)\b|"
    r"\b(não|quando|você|está|sempre|assim|jogo|cada|também|"
    r"porque|entre|antes|depois|mesma|mesmo|outro|outra)\b",
    re.I,
)


def auditar_offline(entradas: list[dict]) -> int:
    problemas: list[str] = []
    vistas: dict[str, str] = {}
    for e in entradas:
        cmd = e["command"]
        if not cmd.rstrip().endswith("%command%"):
            problemas.append(f"não termina em %command%: {cmd!r}")
        if cmd in vistas:
            problemas.append(f"comando duplicado: {cmd!r}")
        vistas[cmd] = e["title"][0]
        for nome in ("title", "category", "compat", "description"):
            for idioma in (0, 1):
                if not e[nome][idioma].strip():
                    problemas.append(
                        f"{nome}.{'PT' if idioma == 0 else 'EN'} vazio em {cmd!r}")
        if not e["compat"][0].strip():
            problemas.append(f"compat vazio em {cmd!r}")
        for nome in ("title", "category", "compat", "description"):
            for idioma, rotulo in ((0, "PT"), (1, "EN")):
                achado = GLITCH.search(e[nome][idioma])
                if achado:
                    problemas.append(
                        f"{nome}.{rotulo} com resto de edição em {cmd!r}: "
                        f"{achado.group()!r}")
        for nome in ("title", "category", "compat", "description"):
            achado = INGLES_NO_PT.search(e[nome][0])
            if achado:
                problemas.append(
                    f"{nome}.PT com verbo inglês em {cmd!r}: {achado.group()!r}")
            achado = PT_NO_EN.search(e[nome][1])
            if achado:
                problemas.append(
                    f"{nome}.EN com palavra portuguesa em {cmd!r}: "
                    f"{achado.group()!r}")
        toks = cmd.split()
        idx_flags = [i for i, t in enumerate(toks) if t.startswith("-") and t != "--"]
        sep = toks.index("--") if "--" in toks else None
        if idx_flags:
            if idx_flags[0] != 1:
                problemas.append(
                    f"wrapper não é o primeiro token: {cmd!r}")
            if sep is not None and any(i > sep for i in idx_flags):
                problemas.append(f"flag depois do --: {cmd!r}")
            if any("=" in toks[i] for i in idx_flags):
                problemas.append(f"env var entre flags: {cmd!r}")
            if any("=" in toks[i] for i, t in enumerate(toks) if i > idx_flags[0] and "=" in t
                   and t not in ("--",) and not t.startswith("-")):
                problemas.append(
                    f"env var depois da flag: o buildCombination só garante env antes de wrapper: {cmd!r}")
        for v in env_vars(cmd):
            if v in IGNORAR:
                problemas.append(f"token que parece env var mas é termo solto: {v} em {cmd!r}")
        valores: dict[str, set[str]] = collections.defaultdict(set)
        for tok in cmd.replace("%command%", " ").split():
            if "=" in tok:
                k, _, val = tok.partition("=")
                if re.fullmatch(r"[A-Z][A-Z0-9_]+", k):
                    valores[k].add(val.strip('"'))
        for k, vs in valores.items():
            if len(vs) > 1:
                problemas.append(
                    f"{k} com valores diferentes no mesmo comando ({', '.join(sorted(vs))}): {cmd!r}")
        depois = cmd.split("%command%", 1)
        if len(depois) > 1 and depois[1].strip():
            problemas.append(f"token depois de %command%: {cmd!r}")

    print(f"catálogo: {len(entradas)} entradas")
    if problemas:
        print(f"\n{len(problemas)} problema(s):")
        for p in problemas:
            print(f"  - {p}")
        return 1
    print("invariantes locais: ok")
    return 0


# --------------------------------------------------------------------------- online

def auditar_online(entradas: list[dict], cache: str) -> int:
    textos = {nome: baixar(nome, url, cache) for nome, url in FONTES.items()}
    faltando = [n for n, t in textos.items() if not t]
    if faltando:
        print(f"fontes indisponíveis: {', '.join(faltando)}\n", file=sys.stderr)

    onde = collections.defaultdict(set)
    for nome, txt in textos.items():
        for m in PADRAO_VAR.finditer(txt):
            if m.group(1) not in IGNORAR:
                onde[m.group(1)].add(nome)

    por_prefixo = collections.defaultdict(set)
    for e in entradas:
        for v in env_vars(e["command"]):
            por_prefixo[v].add(e["command"])

    print(f"catálogo: {len(entradas)} entradas, {len(por_prefixo)} env vars distintas")
    print(f"fontes lidas: {len(textos) - len(faltando)}/{len(textos)}")

    print("\n== env vars sem menção em nenhuma fonte ==")
    suspeitas = sorted(v for v in por_prefixo if v not in onde)
    if not suspeitas:
        print("  (nenhuma)")
    for v in suspeitas:
        onde_ = destino_de(v)
        print(f"  {v}")
        print(f"      conferir em: {onde_}")
        for c in sorted(por_prefixo[v]):
            print(f"      entrada: {c}")

    print("\n== env vars que o upstream removeu ou renomeou ==")
    sem_fonte = sorted(v for v in por_prefixo if v in REMOVIDAS and v not in onde)
    if not sem_fonte:
        print("  (nenhuma)")
    for v in sem_fonte:
        print(f"  {v}: {REMOVIDAS[v]}")
        for c in sorted(por_prefixo[v]):
            print(f"      entrada: {c}")

    print("\n== env vars que o catálogo documenta como removidas (confirme a nota) ==")
    citadas = sorted(v for v in por_prefixo if v in REMOVIDAS and v in onde)
    for v in citadas:
        print(f"  {v} -> {REMOVIDAS[v]}")
    if not citadas:
        print("  (nenhuma)")

    return 0


def main() -> int:
    ap = argparse.ArgumentParser(description=__doc__,
                                 formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--offline", action="store_true", help="só invariantes locais")
    ap.add_argument("--online", action="store_true", help="cruzar com as fontes upstream")
    ap.add_argument("--cache", default=CACHE, help=f"diretório de cache (padrão: {CACHE})")
    args = ap.parse_args()

    if not os.path.exists(CATALOGO):
        print(f"erro: {CATALOGO} não encontrado; rode a partir do repositório", file=sys.stderr)
        return 2

    entradas = entradas_do_catalogo()
    rc = 0
    if args.offline or not args.online:
        rc |= auditar_offline(entradas)
    if args.online:
        rc |= auditar_online(entradas, args.cache)
    return rc


if __name__ == "__main__":
    sys.exit(main())
