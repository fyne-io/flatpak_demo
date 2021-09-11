# flatpak_demo
A demo of a Fyne application packaged as a Flatpak. This project is intended to serve as a helpful starting point to help developers package their Fyne applications using Flatpak.

## Requirements
Both `flatpak` and `flatpak-builder` need to be installed in order to build the packages. Commands for installing Flatpak can be found [here](https://flatpak.org/setup/). Installing the builder should be very similar.
More information about how to build Flatpaks can be found [here](https://docs.flatpak.org/en/latest/first-build.html) and [here](https://docs.flatpak.org/en/latest/building.html).

With that installed, we need to install the development SDKs that we are going to use:
```
flatpak install flathub org.freedesktop.Platform//21.08 org.freedesktop.Sdk//21.08 org.freedesktop.Sdk.Extension.golang//21.08
```

## Building and installing
Building the package can be done with the following command:
```
flatpak-builder --force-clean build-dir io.fyne.flatpak_demo.yml
```

The package can also be built and installed in one step, using the following command:
```
flatpak-builder --user --install --force-clean build-dir io.fyne.flatpak_demo.yml
```

## Vendoring
It is possible to enable network access during the build, but it is recommended to not do so. If you want to publish the package to [flathub](https://flathub.org), it is even a requirement.
For that reason, it is a good idea to vendor the project using `go mod vendor` and add the files to the git repo.

## Sandbox permissions
The Flatpak applications run within a sandbox that restricts their communication with the host system. It is generally preffered to have the application be as strict as possible.
Our example app only opens up filesystem access for the `Documents` folder and the Fyne preferences and document storage system. Remove these if your app do not need it.
More information on avaliable permissions can be found [here](https://docs.flatpak.org/en/latest/sandbox-permissions.html).

## Metainfo
Flatpak requires the application to provide metadata about itself. This metadata can easily be created using the [AppStream Metainfo Creator](https://www.freedesktop.org/software/appstream/metainfocreator/#/guiapp).

## Example manifest
The manifest for this project can be found below. It can be used as a base to use when packaging other apps using Flatpak.

```yml
app-id: io.fyne.flatpak_demo # Needs to be the same as the Fyne app-id.
runtime: org.freedesktop.Platform
runtime-version: '21.08'
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

    # Open up the documents folder as a minimal example.
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
          sha256: f57aac7463a8dee7884d5c9873277dc71b65dc9fe0a5f2d6636ff9de58be9008
```
