// Package main provides various examples of Fyne API capabilities
package main

import "errors"
import "fmt"
import "log"

import "github.com/fyne-io/fyne/examples/apps"

import "github.com/fyne-io/fyne"
import "github.com/fyne-io/fyne/layout"
import "github.com/fyne-io/fyne/theme"
import "github.com/fyne-io/fyne/dialog"
import W "github.com/fyne-io/fyne/widget"

func canvasApp(app fyne.App) {
	apps.Canvas(app)
}

func layoutApp(app fyne.App) {
	apps.Layout(app)
}

func appButton(app fyne.App, label string, onClick func(fyne.App)) *W.Button {
	return &W.Button{Text: label, OnTapped: func() {
		onClick(app)
	}}
}

func confirmCallback(response bool) {
	log.Println("Responded with", response)
}

func main() {
	app := apps.NewApp()

	w := app.NewWindow("Examples")
	w.SetContent(&W.Box{Children: []fyne.CanvasObject{
		&W.Label{Text: "Fyne Examples!"},

		&W.Button{Text: "Apps", OnTapped: func() {
			dialog.ShowInformation("Information", "Example applications have moved to https://github.com/fyne-io/examples", app)
		}},

		W.NewGroup("Demos", []fyne.CanvasObject{
			appButton(app, "Canvas", canvasApp),
			appButton(app, "Layout", layoutApp),
			&W.Entry{Text: "Entry"},
			&W.Check{Text: "Check", OnChanged: func(on bool) { fmt.Println("checked", on) }},
		}...),

		W.NewGroup("Dialogs", []fyne.CanvasObject{
			&W.Button{Text: "Info", OnTapped: func() {
				dialog.ShowInformation("Information", "You should know this thing...", app)
			}},
			&W.Button{Text: "Error", OnTapped: func() {
				err := errors.New("A dummy error message")
				dialog.ShowError(err, app)
			}},
			&W.Button{Text: "Confirm", OnTapped: func() {
				dialog.ShowConfirm("Confirmation", "Do you want to confirm?", confirmCallback, app)
			}},
			&W.Button{Text: "Custom", OnTapped: func() {
				dialog.ShowCustom("Custom Dialog", &W.Check{Text: "Inside a dialog"}, app)
			}},
		}...),
		layout.NewSpacer(),

		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			&W.Button{Text: "Dark", OnTapped: func() {
				fyne.GetSettings().SetTheme("dark")
			}},
			&W.Button{Text: "Light", OnTapped: func() {
				fyne.GetSettings().SetTheme("light")
			}},
		),
		&W.Button{Text: "Quit", Icon: theme.CancelIcon(), OnTapped: func() {
			app.Quit()
		}},
	}})
	w.Show()
}
