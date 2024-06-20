package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageMainTableType struct{}

var pageMainTable pageMainTableType
var tableQueryResults *tview.Table

func (pageMainTable *pageMainTableType) build() {

	tableQueryResults = tview.NewTable()
	tableQueryResults.Box.SetBorder(true)
	tableQueryResults.Box.SetTitle("Results (alt+r)")
	tableQueryResults.SetBorders(true)
	tableQueryResults.SetSelectable(false, false)
	tableQueryResults.SetBackgroundColor(tcell.ColorBlack)

	tableQueryResults.SetFocusFunc(func() {
		tableQueryResults.SetSelectable(true, false)
	})

	tableQueryResults.SetBlurFunc(func() {
		tableQueryResults.SetSelectable(false, false)
	})

	pagesMain.AddPage("tableQueryResults", tableQueryResults, true, false)
}
