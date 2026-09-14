#!/bin/bash
cd "$(dirname "$0")" || exit 1
set -euo pipefail
source packaging/common.sh

build_binary

echo "==> Gerando .rpm..."
if ! command -v rpmbuild >/dev/null 2>&1; then
  echo "rpmbuild não encontrado. Instale com: sudo dnf install rpm-build (Fedora) ou sudo apt install rpm (Debian/Ubuntu)" >&2
  exit 1
fi

RPM_DIR="build/rpm"
SRCTREE="$RPM_DIR/srctree/${NAME}-${VERSION}"
rm -rf "$RPM_DIR"
mkdir -p "$RPM_DIR"/{BUILD,RPMS,SOURCES,SPECS,SRPMS}
mkdir -p "$SRCTREE"
cp "$NAME" "$SRCTREE/"
cp FyneApp.toml "$SRCTREE/FyneApp.toml"
cp "$DESKTOP" "$SRCTREE/protoncommand.desktop"
cp "$ICON" "$SRCTREE/icon.png"
tar -C "$RPM_DIR/srctree" -czf "$RPM_DIR/SOURCES/${NAME}-${VERSION}.tar.gz" "${NAME}-${VERSION}"
cp packaging/protoncommand.spec "$RPM_DIR/SPECS/"
sed -i "s/^Version: .*/Version: $VERSION/" "$RPM_DIR/SPECS/protoncommand.spec"
rpmbuild --define "_topdir $(pwd)/$RPM_DIR" -bb "$RPM_DIR/SPECS/protoncommand.spec"
mkdir -p "$DIST"
cp "$RPM_DIR"/RPMS/x86_64/*.rpm "$DIST/${ARTIFACT}.rpm"
echo "==> .rpm pronto em $DIST/"
