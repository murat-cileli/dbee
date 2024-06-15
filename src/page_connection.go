package main

import (
	"github.com/rivo/tview"
)

var formConnectionNew *tview.Form
var listConnections *tview.List
var flexConnection *tview.Flex

type pageConnectionType struct{}

func (pageConnection *pageConnectionType) build() *pageConnectionType {
	listConnections = tview.NewList()
	listConnections.SetBorder(true).SetTitle("Saved Connections").SetTitleAlign(tview.AlignCenter)

	formConnectionNew = tview.NewForm().
		AddDropDown("Driver", []string{"MySQL"}, 0, nil).
		AddInputField("Connection (*)", "root@tcp(localhost:3306)/mydb", 0, nil, nil).
		AddPasswordField("Password", "root", 0, '*', nil).
		AddButton("Connect", func() {
			_, database.Driver = formConnectionNew.GetFormItemByLabel("Driver").(*tview.DropDown).GetCurrentOption()
			database.ConnectionString = formConnectionNew.GetFormItemByLabel("Connection (*)").(*tview.InputField).GetText()
			database.Password = formConnectionNew.GetFormItemByLabel("Password").(*tview.InputField).GetText()
			err := database.Connect()
			if err != nil {
				pageAlert.show(err.Error(), "error")
				return
			} else {
				if formConnectionNew.GetFormItemByLabel("Save connection").(*tview.Checkbox).IsChecked() {
					application.saveConnection(database.Driver, database.ConnectionString)
				}
				pageMain.build()
				pages.ShowPage("main")
				app.SetFocus(textAreaQuery)
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		}).
		AddCheckbox("Save connection", false, nil)

	formConnectionNew.SetBorder(true).SetTitle("Connect to Server").SetTitleAlign(tview.AlignCenter)
	formConnectionNew.SetButtonsAlign(tview.AlignCenter)

	flexConnection = tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(listConnections, 0, 3, false).
		AddItem(formConnectionNew, 0, 5, true).
		AddItem(nil, 0, 1, false)

	pages.AddPage("connection", flexConnection, true, false)

	return pageConnection
}

func (pageConnection *pageConnectionType) show() {
	pages.ShowPage("connection")
}
