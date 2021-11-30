package main

import (
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("io.fyne.flatpak_demo")
	w := a.NewWindow("Flatpak Demo")

	markdown := widget.NewMultiLineEntry()
	preview := widget.NewRichText()
	markdown.OnChanged = preview.ParseMarkdown

	open := &widget.Button{Text: "Open file", Icon: theme.ContentAddIcon(), OnTapped: func() {
		files := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println(err)
				return
			} else if file == nil {
				return
			}

			text, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println(err)
				return
			}

			markdown.SetText(string(text))
			a.SendNotification(&fyne.Notification{Title: "Opened a file", Content: file.URI().Name() + " was opened correctly."})
		}, w)

		files.SetFilter(storage.NewExtensionFileFilter([]string{".md"}))
		files.Show()
	}}

	w.SetContent(
		container.NewBorder(container.NewHBox(open), nil, nil, nil,
			container.NewHSplit(markdown, preview),
		),
	)

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
