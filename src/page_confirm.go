package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageConfirmType struct {
	*tview.Modal
}

var pageConfirm pageConfirmType

func (pageConfirm *pageConfirmType) build() {
	pageConfirm.Modal = tview.NewModal().AddButtons([]string{"Yes", "No"})
	pageConfirm.Modal.SetBorder(true).SetTitleAlign(tview.AlignCenter)
	pageConfirm.Modal.Box.SetBackgroundColor(tcell.ColorDarkBlue.TrueColor())
	pageConfirm.Modal.SetBackgroundColor(tcell.ColorDarkBlue.TrueColor())
	pageConfirm.Modal.SetTitle("Confirmation")
	pageConfirm.Modal.Box.SetBorderColor(tcell.ColorWhite.TrueColor())
	pageConfirm.SetTitleColor(tcell.ColorWhite.TrueColor())
	pageConfirm.Modal.SetTextColor(tcell.ColorWhite.TrueColor())
	application.pages.AddPage("confirm", pageConfirm.Modal, true, false)
}

func (pageConfirm *pageConfirmType) show(message string, callback func()) {
	pageConfirm.Modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Yes" {
			callback()
		}
		application.pages.HidePage("confirm")
	})
	pageConfirm.Modal.SetText(message)
	application.pages.ShowPage("confirm").SendToFront("confirm")
}
