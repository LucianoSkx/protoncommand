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
- Filtro por categoria e favoritos; busca por comando, título, descrição, categoria e compatibilidade
- Tema claro / escuro / sistema
- Opção "copiar ao clicar"
- Seletor de idioma (PT/EN) no topo

## Como rodar

Com o AppImage (ver Instalação abaixo) ou compilando do código-fonte:

```bash
go build -o protoncommand .
./protoncommand
```

## Dependências de compilação

Precisa de Go + `gcc` + bibliotecas do GLFW:

| Distro | Instalar |
|---|---|
| **Arch / CachyOS** | `sudo pacman -S base-devel go` |
| **Debian / Ubuntu** | `sudo apt install golang gcc libgl1-mesa-dev xorg-dev libwayland-dev` |
| **Fedora** | `sudo dnf install golang gcc mesa-libGL-devel libX11-devel libwayland-devel libXcursor-devel libXrandr-devel libXi-devel` |

O binário é sempre gerado dentro do AppImage; não é preciso mais nada.

## Instalação

Baixe o `.AppImage` mais recente na página de [Releases](https://github.com/LucianoSkx/protoncommand/releases):

```bash
chmod +x protoncommand-*.x86_64.AppImage
./protoncommand-*.x86_64.AppImage
```

### Atualizações

O AppImage traz informação de atualização embutida: use [AppImageUpdate](https://github.com/AppImageCommunity/AppImageUpdate), `appimageupdatetool` ou Gear Lever para atualizar só o que mudou (delta) direto dos Releases do GitHub, sem baixar tudo de novo.

## Compilando o AppImage

```bash
./build-appimage.sh         # usa a tag git mais recente (vX.Y.Z)
./build-appimage.sh 0.5.0   # ou informe a versão
```

O AppImage + `.zsync` saem em `dist/`.

## Contribuindo

```bash
go test ./...   # testes (catálogo, combinação, conflitos, i18n)
go vet ./...    # análise estática
gofmt -l .      # deve sair vazio
```

Todo comando novo em `commands.go` precisa de título, categoria, compatibilidade e descrição em PT e EN, e terminar com `%command%`. O CI valida unicidade, paridade de idiomas e a lista de variáveis obsoletas (`TestNoObsoleteCommands`).

Para lançar uma versão: `git tag vX.Y.Z && git push origin vX.Y.Z` — o GitHub gera o AppImage e publica o Release sozinho.

## Licença

MIT — veja [LICENSE](LICENSE).

---

# English 🇺🇸

Manage and combine useful Proton launch commands (standard Proton, Proton-GE and Proton-CachyOS) in a simple GUI app. Click a command to see what it does, check several to build a combination ready to paste into Steam launch options.

## Features

- **79 commands** with descriptions in English and Portuguese
- **Multiple combination**: check several commands and generate a single ready-to-paste line, e.g.:
  `mangohud gamemoderun PROTON_LOG=1 %command%`
- **Launcher selector**: Steam, Faugus Launcher, Heroic, Lutris and Bottles — the generated line adapts (Steam/Faugus use `%command%`; the rest get environment variables only)
- **Conflict warnings**: detects variables set more than once with different values and mutually exclusive options (e.g.: Anti-Lag 2 vs Reflex)
- **Favorites**: star commands and filter favorites only
- Category and favorites filter; search by command, title, description, category and compatibility
- Light / dark / system theme
- Optional "copy on click"
- Language selector (PT/EN) at the top

## Running

With the AppImage (see Installation below) or building from source:

```bash
go build -o protoncommand .
./protoncommand
```

## Build dependencies

You need Go + `gcc` + the GLFW libraries:

| Distro | Install |
|---|---|
| **Arch / CachyOS** | `sudo pacman -S base-devel go` |
| **Debian / Ubuntu** | `sudo apt install golang gcc libgl1-mesa-dev xorg-dev libwayland-dev` |
| **Fedora** | `sudo dnf install golang gcc mesa-libGL-devel libX11-devel libwayland-devel libXcursor-devel libXrandr-devel libXi-devel` |

The binary is always built inside the AppImage; nothing else is needed.

## Installation

Download the latest `.AppImage` from the [Releases](https://github.com/LucianoSkx/protoncommand/releases) page:

```bash
chmod +x protoncommand-*.x86_64.AppImage
./protoncommand-*.x86_64.AppImage
```

### Updates

The AppImage ships embedded update information: use [AppImageUpdate](https://github.com/AppImageCommunity/AppImageUpdate), `appimageupdatetool` or Gear Lever to apply delta updates straight from the GitHub Releases, without downloading everything again.

## Building the AppImage

```bash
./build-appimage.sh         # uses the latest git tag (vX.Y.Z)
./build-appimage.sh 0.5.0   # or pass the version
```

The AppImage + `.zsync` land in `dist/`.

## Contributing

```bash
go test ./...   # tests (catalog, combination, conflicts, i18n)
go vet ./...    # static analysis
gofmt -l .      # must print nothing
```

Every new command in `commands.go` needs title, category, compatibility and description in both EN and PT, and must end with `%command%`. CI checks uniqueness, language parity and the obsolete-variable list (`TestNoObsoleteCommands`).

To cut a release: `git tag vX.Y.Z && git push origin vX.Y.Z` — GitHub builds the AppImage and publishes the Release on its own.

## License

MIT — see [LICENSE](LICENSE).
