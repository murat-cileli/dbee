package main

import (
	"github.com/rivo/tview"
)

var formConnectionNew *tview.Form

type pageConnectionNewType struct{}

func (pageConnectionNew *pageConnectionNewType) build() *pageConnectionNewType {
	formConnectionNew = tview.NewForm().
		AddDropDown("Driver", []string{"MariaDB / MySQL"}, 0, nil).
		AddInputField("Host (*)", "localhost", 0, nil, nil).
		AddInputField("Port (*)", "3306", 10, nil, nil).
		AddInputField("User (*)", "root", 0, nil, nil).
		AddPasswordField("Password", "root", 0, '*', nil).
		AddInputField("Database", "mydb", 0, nil, nil).
		AddButton("Connect", func() {
			database.Host = formConnectionNew.GetFormItemByLabel("Host (*)").(*tview.InputField).GetText()
			database.Port = formConnectionNew.GetFormItemByLabel("Port (*)").(*tview.InputField).GetText()
			database.User = formConnectionNew.GetFormItemByLabel("User (*)").(*tview.InputField).GetText()
			database.Password = formConnectionNew.GetFormItemByLabel("Password").(*tview.InputField).GetText()
			database.Database = formConnectionNew.GetFormItemByLabel("Database").(*tview.InputField).GetText()
			err := database.Connect()
			if err != nil {
				pageAlert.show(err.Error(), "error")
				return
			} else {
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

	flexConnectionNew := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(formConnectionNew, 0, 2, true).
		AddItem(nil, 0, 1, false)

	pages.AddPage("connectionNew", flexConnectionNew, true, false)

	return pageConnectionNew
}

func (pageConnectionNew *pageConnectionNewType) show() {
	pages.ShowPage("connectionNew")
}
