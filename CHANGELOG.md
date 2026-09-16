# v0.5.5

Favoritos mais robustos e cobertura de testes quase triplicada.

## 🐛 Correções

- Favoritos usam chave composta (comando + título) para evitar colisão quando dois comandos diferentes têm o mesmo shell.

## 🛠️ Por baixo do pano

- Cobertura de testes subiu de 19.8% para 38.7%.
- Novos testes: toggleFav, clearSelection, selectCommand, clearDetail, updateFavButton, setTheme, setLang, launcher, updateCombination, applyFilter, combinação, conflitos, splitFields, displayCmd, tr, t, cmd.

# v0.5.4

Detalhe acompanha a lista ao pesquisar e ao trocar de idioma.

## 🐛 Correções / Fixes

- Novo `maxWidthLabel`: combinação, status e detalhes quebram linha dentro da janela em vez de esticá-la além da tela.
- Removida a rolagem (`VScroll`) da barra de combinação.
- Pesquisa: painel de detalhes agora acompanha a lista mesmo quando a linha 0 já estava selecionada.
- Idioma: detalhe atualiza o texto ao trocar PT/EN (mesmo bug do search).

# v0.5.3

Aparece só na aba Jogos do menu do sistema.

## 🖥️ Menu

- `Categories=Game;` no `.desktop` (removido `Utility`) — o app não aparece mais em Utilitários.

# v0.5.2

lsfg-vk atualizado para v2.0.0+ e barra de combinação com rolagem.

## 📚 Catálogo (78 comandos)

- **Removido**: `LSFG_PROCESS=steam` (variável removida no lsfg-vk v2.0.0).

## 🐛 Correções

- Barra de combinação com rolagem (`NewVScroll`) — não expande mais a janela além da tela com muitos comandos.
- README com exemplo de combinação e versão atualizados.

# v0.5.1

Ordem correta dos wrappers na combinação e lsfg-vk v2.0.0+.

## 📚 Catálogo

- **Novo**: `LSFGVK_PROFILE=steam` (variável oficial do lsfg-vk v2.0.0+ para selecionar o perfil).

## 🐛 Correções

- Wrappers (`game-performance`, `mangohud`, `gamemoderun`, `gamescope`) agora vão por último, antes do `%command%`: `PROTON_LOG=1 mangohud gamemoderun %command%`.
- Label da combinação contido (`maxWidthLabel`) para não esticar a janela.

# v0.5.0

Catálogo auditado, AppImage como formato único (com atualização delta) e releases automáticos pelo GitHub.

## 📚 Catálogo (78 comandos)

- **Removidos 18**: variáveis obsoletas, renomeadas ou sem efeito nas versões atuais — `DXVK_ASYNC`, `PROTON_USE_NTSYNC`, `PROTON_ENABLE_NVAPI`, `PROTON_VKREFLEX`, `PROTON_VKD3D_HEAP`, `PROTON_NO_D3D9`, `PROTON_ENABLE_HDR`, `PROTON_DUMP_DEBUG_COMMANDS`, `WINE_AUDIO_DRIVER`, `WINE_BLOCK_HOSTS`, `WINE_VIRTUAL_DESKTOP`, `WINE_ESYNC`, `WINEFSYNC`, `FNA3D_FORCE_DRIVER`, `DRI_CONFIG` — e as 3 entradas de cracks (Online-Fix/SteamFix).
- **Novos 4**: `PROTON_FFX3_UPGRADE` (nome atual do FSR 3.1 no CachyOS 11+), `PROTON_VKD3D_LOWLATENCY` (D3D12), `DXVK_FRAME_RATE=60` (limite de FPS sem overlay), `PROTON_D7VK_DDRAW` (jogos DX7 ou anteriores).
- **Correções**: `DXVK_HUD` (`gpu` → `gpuload`); escopo explícito DX8–11 vs DX12 no low-latency; compatibilidades revisadas (RDNA3, indicadores FSR4/DLSS, LSFG).
- Removida a duplicata `PROTON_XESS_UPGRADE`.

## 📦 Distribuição

- Só AppImage: removidos `.deb`, `.rpm` e scripts por formato.
- Nome com versão (`protoncommand-0.5.0-x86_64.AppImage`) e update info embutida (zsync) — atualize por delta com AppImageUpdate ou Gear Lever.
- Release publicado automaticamente pelo GitHub a cada tag `v*`.

## 🛠️ Por baixo do pano

- Combinação: wrappers vão por último, antes de `%command%` (ordem correta: env vars → wrappers → %command%).
- Detector de conflitos respeita aspas e compara nomes exatos.
- CI em todo push/PR (`gofmt`, `vet`, `test`, `build`) com retry de download e fallback de proxy; testes anti-regressão do catálogo e i18n.

## 🛠️ Instalação rápida

```bash
chmod +x protoncommand-0.5.0-x86_64.AppImage && ./protoncommand-0.5.0-x86_64.AppImage
```

# v0.4.0

Correções de nomenclatura e internacionalização do rótulo de idioma.

## 🐛 Correções

- **Licença**: copyright atualizado de `ProtonBox contributors` para `protoncommand contributors` no `LICENSE`.
- **README**: seção em Português agora diz "descrição em Português e Inglês" (antes "English").
- **App**: o seletor de idioma e o menu de configurações passam a exibir **Inglês** em vez de "English" (PT e EN).

## 📦 Pacotes

| Formato | Arquivo | Uso |
|---|---|---|
| AppImage | `protoncommand-0.4.0.x86_64.AppImage` | qualquer distro |
| Debian/Ubuntu | `protoncommand-0.4.0.x86_64.deb` | `sudo dpkg -i` |
| Fedora/openSUSE | `protoncommand-0.4.0.x86_64.rpm` | `sudo rpm -i` |

## 🛠️ Instalação rápida

```bash
# AppImage
chmod +x protoncommand-0.4.0.x86_64.AppImage && ./protoncommand-0.4.0.x86_64.AppImage

# Debian/Ubuntu
sudo dpkg -i protoncommand-0.4.0.x86_64.deb

# Fedora/openSUSE
sudo rpm -i protoncommand-0.4.0.x86_64.rpm
```

# v0.3.0

Catálogo expandido e descrições reescritas com foco prático (sintoma → solução, exemplos de valores e consistência PT/EN).

## ✨ Novos comandos

### Upscalers e upgrades
- `PROTON_FSR4_INDICATOR=1`, `PROTON_FSR4_RDNA3_UPGRADE=1`, `PROTON_FSR3_UPGRADE=1` — indicador e upgrades FSR4/FSR3
- `PROTON_XESS_UPGRADE=1` — XeSS (atualizado)
- `PROTON_DLSS_INDICATOR=1` — indicador DLSS
- `PROTON_OPTISCALER_NAME=dxgi.dll`, `PROTON_OPTISCALER_CONFIG=...` — OptiScaler

### Latência e HDR
- `DXVK_NVAPI_VKREFLEX=1`, `PROTON_VKREFLEX=1` — NVIDIA Reflex em jogos Vulkan
- `PROTON_ENABLE_HDR=1` — HDR via Proton
- `PROTON_VKD3D_HEAP=1` — heap dedicada do vkd3d

### Display e input
- `PROTON_USE_WAYLAND=1`, `PROTON_PREFER_SDL=1`, `PROTON_NO_STEAMINPUT=1`

### Diversos
- `PROTON_ADD_CONFIG=config1,config2` — aplica configs do Proton sem editar arquivos
- `PROTON_LOCAL_SHADER_CACHE=1` — cache local de shaders
- `PROTON_MEDIA_FORCE_GST=1`, `PROTON_GST_VIDEO_ORIENTATION=90`, `WINE_BLOCK_HOSTS=...`

### Obsoleto
- `PROTON_USE_NTSYNC=1` — marcado como obsoleto (ntsync é o padrão no Proton 11+; use `PROTON_NO_NTSYNC=1` para desativar)

## 📦 Pacotes

| Formato | Arquivo | Uso |
|---|---|---|
| AppImage | `protoncommand-0.3.0.x86_64.AppImage` | qualquer distro |
| Debian/Ubuntu | `protoncommand-0.3.0.x86_64.deb` | `sudo dpkg -i` |
| Fedora/openSUSE | `protoncommand-0.3.0.x86_64.rpm` | `sudo rpm -i` |

## 🛠️ Instalação rápida

```bash
# AppImage
chmod +x protoncommand-0.3.0.x86_64.AppImage && ./protoncommand-0.3.0.x86_64.AppImage

# Debian/Ubuntu
sudo dpkg -i protoncommand-0.3.0.x86_64.deb

# Fedora/openSUSE
sudo rpm -i protoncommand-0.3.0.x86_64.rpm
```
