# This is just a Makefile to simplify running some of the commands mentioned in the README.

build:
	flatpak-builder --user --force-clean build-dir io.fyne.flatpak_demo.yml

install:
	flatpak-builder --user --install --force-clean build-dir io.fyne.flatpak_demo.yml

run:
	flatpak run --user io.fyne.flatpak_demo
