<div align="center">
  <img src="assets/logo.png" alt="Proton Command" width="400">
</div>

<div align="center">

[![Release](https://img.shields.io/github/v/release/LucianoSkx/protoncommand)](https://github.com/LucianoSkx/protoncommand/releases)
[![CI](https://github.com/LucianoSkx/protoncommand/actions/workflows/ci.yml/badge.svg)](https://github.com/LucianoSkx/protoncommand/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

</div>

---

# Português 🇧🇷

Gerencie e combine comandos úteis do Proton (Proton padrão, Proton-GE e Proton-CachyOS) em um app gráfico simples. Clique em um comando para ver o que ele faz, marque vários para montar uma combinação pronta para colar nas opções de inicialização do Steam.

## Funcionalidades

- **78 comandos** com descrição em Português e Inglês
- **Combinação múltipla**: marque vários comandos e gere uma única linha pronta, por exemplo:
  `mangohud gamemoderun PROTON_LOG=1 %command%`
- **Seletor de launcher**: Steam, Faugus Launcher, Heroic, Lutris e Bottles — a linha gerada se adapta (Steam/Faugus usam `%command%`; os demais só variáveis de ambiente)
- **Avisos de conflito**: detecta variáveis definidas mais de uma vez com valores diferentes e opções mutuamente exclusivas (ex.: Anti-Lag 2 vs Reflex)
- **Favoritos**: marque comandos com estrela e filtre só os favoritos
- Filtro por categoria, busca por comando, título, categoria e compatibilidade
- Tema claro / escuro / sistema
- Opção "copiar ao clicar"
- Seletor de idioma (PT/EN) no topo

## Como rodar

```bash
./protoncommand
```

Ou compile do código-fonte:

```bash
go build -o protoncommand .
```

## Dependências de compilação

Precisa de Go + `gcc` + bibliotecas do GLFW:

| Distro | Instalar |
|---|---|
| **Arch / CachyOS** | `sudo pacman -S base-devel go` |
| **Debian / Ubuntu** | `sudo apt install golang gcc libgl1-mesa-dev xorg-dev libwayland-dev` |
| **Fedora** | `sudo dnf install golang gcc mesa-libGL-devel libX11-devel libwayland-devel libXcursor-devel libXrandr-devel libXi-devel` |

O binário + AppImage sempre são gerados; o `.deb` precisa de `dpkg-dev` e o `.rpm` de `rpm-build`.

## Instalação

Escolha um dos pacotes na página de [Releases](https://github.com/LucianoSkx/protoncommand/releases):

| Formato | Uso |
|---|---|
| **AppImage** | Qualquer distro: `chmod +x protoncommand-0.4.0.x86_64.AppImage && ./protoncommand-0.4.0.x86_64.AppImage` |
| **.deb** | Debian/Ubuntu: `sudo dpkg -i protoncommand-0.4.0.x86_64.deb` |
| **.rpm** | Fedora/openSUSE: `sudo rpm -i protoncommand-0.4.0.x86_64.rpm` |
| **Binário** | Compile com `./build-binary.sh` |

## Compilando os pacotes

Builds separados por formato:

```bash
./build-binary.sh     # binário
./build-appimage.sh   # AppImage (qualquer distro)
./build-deb.sh        # .deb (precisa dpkg-dev)
./build-rpm.sh        # .rpm (precisa rpm-build)
```

Artefatos em `dist/`. Rode o script do formato desejado — cada um compila o binário e gera seu pacote.

## Contribuindo

```bash
go test ./...   # testes (catálogo, combinação, conflitos, i18n)
gofmt -l .      # deve sair vazio
```

Todo comando novo em `commands.go` precisa de título, categoria, compatibilidade e descrição em PT e EN, e terminar com `%command%`. O CI valida unicidade e paridade de idiomas.

## Licença

MIT — veja [LICENSE](LICENSE).

---

# English 🇺🇸

Manage and combine useful Proton launch commands (standard Proton, Proton-GE and Proton-CachyOS) in a simple GUI app. Click a command to see what it does, check several to build a combination ready to paste into Steam launch options.

## Features

- **78 commands** with descriptions in English and Portuguese
- **Multiple combination**: check several commands and generate a single ready-to-paste line, e.g.:
  `mangohud gamemoderun PROTON_LOG=1 %command%`
- **Launcher selector**: Steam, Faugus Launcher, Heroic, Lutris and Bottles — the generated line adapts (Steam/Faugus use `%command%`; the rest get environment variables only)
- **Conflict warnings**: detects variables set more than once with different values and mutually exclusive options (e.g.: Anti-Lag 2 vs Reflex)
- **Favorites**: star commands and filter favorites only
- Category filter, search by command, title, category and compatibility
- Light / dark / system theme
- Optional "copy on click"
- Language selector (PT/EN) at the top

## Running

```bash
./protoncommand
```

Or build from source:

```bash
go build -o protoncommand .
```

## Build dependencies

You need Go + `gcc` + the GLFW libraries:

| Distro | Install |
|---|---|
| **Arch / CachyOS** | `sudo pacman -S base-devel go` |
| **Debian / Ubuntu** | `sudo apt install golang gcc libgl1-mesa-dev xorg-dev libwayland-dev` |
| **Fedora** | `sudo dnf install golang gcc mesa-libGL-devel libX11-devel libwayland-devel libXcursor-devel libXrandr-devel libXi-devel` |

The binary + AppImage are always produced; `.deb` additionally needs `dpkg-dev` and `.rpm` needs `rpm-build`.

## Installation

Pick a package from the [Releases](https://github.com/LucianoSkx/protoncommand/releases) page:

| Format | Usage |
|---|---|
| **AppImage** | Any distro: `chmod +x protoncommand-0.4.0.x86_64.AppImage && ./protoncommand-0.4.0.x86_64.AppImage` |
| **.deb** | Debian/Ubuntu: `sudo dpkg -i protoncommand-0.4.0.x86_64.deb` |
| **.rpm** | Fedora/openSUSE: `sudo rpm -i protoncommand-0.4.0.x86_64.rpm` |
| **Binary** | Build with `./build-binary.sh` |

## Building packages

Separate scripts per format:

```bash
./build-binary.sh     # binary
./build-appimage.sh   # AppImage (any distro)
./build-deb.sh        # .deb (needs dpkg-dev)
./build-rpm.sh        # .rpm (needs rpm-build)
```

Artifacts land in `dist/`. Run the script of the format you want — each one compiles the binary and builds its package.

## Contributing

```bash
go test ./...   # tests (catalog, combination, conflicts, i18n)
gofmt -l .      # must print nothing
```

Every new command in `commands.go` needs title, category, compatibility and description in both EN and PT, and must end with `%command%`. CI checks uniqueness and language parity.

## License

MIT — see [LICENSE](LICENSE).
