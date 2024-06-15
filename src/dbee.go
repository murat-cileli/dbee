package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var pages *tview.Pages
var pageConnectionNew pageConnectionNewType
var pageMain pageMainType
var pageAlert pageAlertType

var database databaseType

func main() {
	app = tview.NewApplication()
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorNone.TrueColor()

	pages = tview.NewPages()
	pageAlert.build()
	pageConnectionNew.build().show()

	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
