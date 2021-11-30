# This is just a simple makefile to simplify running some of the commands mentioned in the README.

build:
	flatpak-builder --force-clean build-dir io.fyne.flatpak_demo.yml

install:
	flatpak-builder --user --install --force-clean build-dir io.fyne.flatpak_demo.yml

run:
	flatpak run io.fyne.flatpak_demo
