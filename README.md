# flatpak_demo
A demo of a Fyne application packaged as a Flatpak. This project is intended to serve as a helpful starting point to help developers package their Fyne applications using Flatpak.

## Requirements
Both `flatpak` and `flatpak-builder` need to be installed in order to build the packages. Commands for installing Flatpak can be found [here](https://flatpak.org/setup/). Installing the builder should be very similar.

With that installed, we need to install the development SDKs that we are going to use:
```
flatpak install flathub org.freedesktop.Platform//21.08 org.freedesktop.Sdk//21.08 org.freedesktop.Sdk.Extension.golang//21.08
```

## Building and installing
Building the command can be done with the following command:
```
flatpak-builder --force-clean build-dir io.fyne.flatpak_demo.yml
```

The package can be built and installed using the following command:
```
flatpak-builder --user --install --force-clean build-dir io.fyne.flatpak_demo.yml
```

## Example manifest

```yml
app-id: io.fyne.flatpak_demo # CHange to your own app-id. Needs to be the same as the Fyne app-id.
runtime: org.freedesktop.Platform
runtime-version: '21.08' # Use the latest version if possible.
sdk: org.freedesktop.Sdk
sdk-extensions:
    - org.freedesktop.Sdk.Extension.golang
command: flatpak_demo

finish-args:
    - --share=ipc # Share IPC namespace with the host (necessary for X11).
    - --socket=x11
    - --device=dri # OpenGL rendering support.

    # Only needed if building with -tags wayland.
    #- --socket=wayland

    # Open up the documents folder as a minimal example. Remove if you don't need filesystem access.
    - --filesystem=xdg-documents

    # Needed for Fyne preferences and document storage.
    - --filesystem=~/.config/fyne/$FLATPAK_ID:create

build-options:
  env:
    - GOBIN=/app/bin
    - GOROOT=/usr/lib/sdk/golang

modules:
    - name: flatpak_demo
      buildsystem: simple
      build-commands:
        - $GOROOT/bin/go build -o flatpak_demo
        - install -Dm00755 flatpak_demo $FLATPAK_DEST/bin/flatpak_demo
        - install -Dm00644 Icon.png $FLATPAK_DEST/share/icons/hicolor/256x256/apps/$FLATPAK_ID.png
        - install -Dm00644 $FLATPAK_ID.desktop $FLATPAK_DEST/share/applications/$FLATPAK_ID.desktop
        - install -Dm00644 $FLATPAK_ID.metainfo.xml $FLATPAK_DEST/share/metainfo/$FLATPAK_ID.metainfo.xml
      sources:
        - type: archive
          url: "https://github.com/Jacalz/flatpak_demo/archive/refs/tags/v1.0.0.tar.gz"
          sha256: #TODO
```
