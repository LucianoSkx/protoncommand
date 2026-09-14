#!/bin/bash
# Gera o AppImage (único formato distribuído).
# Versão: ./build-appimage.sh [versão] | VERSION=x.y.z ./build-appimage.sh
# Sem argumento, usa a tag git mais recente (vX.Y.Z).
cd "$(dirname "$0")" || exit 1
set -euo pipefail

NAME="protoncommand"
ARCH="x86_64"
DESKTOP="assets/protoncommand.desktop"
ICON="assets/icon.png"
DIST="dist"

VERSION="${VERSION:-${1:-}}"
if [ -z "$VERSION" ]; then
  VERSION="$(git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//')"
fi
VERSION="${VERSION:-0.0.0-dev}"
VERSION="${VERSION#v}"
ARTIFACT="${NAME}-${VERSION}-${ARCH}.AppImage"
# Formato canônico com pipes (a forma gh-releases-zsync:// é rejeitada
# pelo appimagetool atual). O * cobre o número da versão no nome.
UPDATE="gh-releases-zsync|LucianoSkx|protoncommand|latest|${NAME}-*-${ARCH}.AppImage.zsync"

echo "==> Compilando ($VERSION)..."
go build -trimpath -ldflags "-s -w" -o "$NAME" .

echo "==> Montando AppDir..."
rm -rf AppDir
mkdir -p AppDir/usr/bin
mkdir -p AppDir/usr/share/applications
mkdir -p AppDir/usr/share/icons/hicolor/256x256/apps
cp "$NAME" AppDir/usr/bin/
cp FyneApp.toml AppDir/usr/bin/FyneApp.toml
cp "$DESKTOP" AppDir/protoncommand.desktop
cp "$DESKTOP" AppDir/usr/share/applications/protoncommand.desktop
cp "$ICON" AppDir/protoncommand.png
cp "$ICON" AppDir/usr/share/icons/hicolor/256x256/apps/protoncommand.png
cat > AppDir/AppRun <<'EOF'
#!/bin/sh
SELF=$(readlink -f "$0")
HERE=${SELF%/*}
exec "$HERE/usr/bin/protoncommand" "$@"
EOF
chmod +x AppDir/AppRun

if command -v appimagetool >/dev/null 2>&1; then
  TOOL="appimagetool"
elif [ -x ./appimagetool ]; then
  TOOL="./appimagetool"
else
  echo "appimagetool não encontrado. Baixando..."
  curl -L -o appimagetool \
    https://github.com/AppImage/appimagetool/releases/download/continuous/appimagetool-x86_64.AppImage
  chmod +x appimagetool
  TOOL="./appimagetool"
fi

mkdir -p "$DIST"
echo "==> Gerando AppImage com update info (zsync)..."
APPIMAGE_EXTRACT_AND_RUN=1 "$TOOL" -u "$UPDATE" AppDir "$DIST/$ARTIFACT"
# O .zsync sai no diretório atual com o nome-base do AppImage; move para dist/.
mv -f "${ARTIFACT}.zsync" "$DIST/" 2>/dev/null || true
echo "==> Pronto: $DIST/$ARTIFACT (+ .zsync para atualizações delta)"
