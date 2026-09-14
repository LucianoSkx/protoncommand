%global debug_package %{nil}

Name: protoncommand
Version: 0.4.0
Release: 1%{?dist}
Summary: Useful Proton launch commands manager
License: MIT
URL: https://github.com/LucianoSkx/protoncommand
Source0: %{name}-%{version}.tar.gz
BuildArch: x86_64

%description
Proton Command is a simple GUI to browse, copy and combine useful Proton launch
commands (standard Proton, Proton-GE and Proton-CachyOS) for Steam.

%prep
%setup -q

%install
mkdir -p %{buildroot}%{_bindir}
install -m755 protoncommand %{buildroot}%{_bindir}/protoncommand
install -m644 FyneApp.toml %{buildroot}%{_bindir}/FyneApp.toml
mkdir -p %{buildroot}%{_datadir}/applications
install -m644 protoncommand.desktop %{buildroot}%{_datadir}/applications/protoncommand.desktop
mkdir -p %{buildroot}%{_datadir}/icons/hicolor/256x256/apps
install -m644 icon.png %{buildroot}%{_datadir}/icons/hicolor/256x256/apps/protoncommand.png

%files
%{_bindir}/protoncommand
%{_bindir}/FyneApp.toml
%{_datadir}/applications/protoncommand.desktop
%{_datadir}/icons/hicolor/256x256/apps/protoncommand.png
