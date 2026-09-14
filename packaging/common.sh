#!/bin/bash
set -euo pipefail

NAME="protoncommand"
# Versão: usa $VERSION do ambiente, senão a tag git mais recente (vX.Y.Z),
# senão o fallback abaixo. Evita artefatos com versão desatualizada.
if [ -z "${VERSION:-}" ]; then
  VERSION="$(git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//')"
fi
VERSION="${VERSION:-0.4.0}"
DESKTOP="packaging/protoncommand.desktop"
ICON="assets/icon.png"
DIST="dist"
ARTIFACT="${NAME}-${VERSION}.x86_64"

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

build_binary() {
  echo "==> Compilando binário..."
  go build -trimpath -ldflags "-s -w" -o "$NAME" .
}

check_binary() {
  if [ ! -x "$NAME" ]; then
    echo "Binário $NAME não encontrado. Execute build-binary.sh antes." >&2
    exit 1
  fi
}
