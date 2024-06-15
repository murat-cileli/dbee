package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageAlertType struct {
	*tview.Modal
}

func (pageAlert *pageAlertType) build() {
	pageAlert.Modal = tview.NewModal().AddButtons([]string{"OK"})
	pageAlert.Modal.SetBorder(true).SetTitleAlign(tview.AlignCenter)
	pageAlert.Modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.HidePage("alert")
	})
	pages.AddPage("alert", pageAlert.Modal, true, false)
}

func (pageAlert *pageAlertType) show(message string, alertType string) {
	if alertType == "error" {
		pageAlert.Modal.Box.SetBackgroundColor(tcell.ColorDarkRed)
		pageAlert.Modal.SetBackgroundColor(tcell.ColorDarkRed)
		pageAlert.Modal.SetTitle("Error")
	} else {
		pageAlert.Modal.Box.SetBackgroundColor(tcell.ColorDarkBlue)
		pageAlert.Modal.SetBackgroundColor(tcell.ColorDarkBlue)
		pageAlert.Modal.SetTitle("Info")
	}
	pageAlert.Modal.SetText(message)
	pages.ShowPage("alert").SendToFront("alert")
}
