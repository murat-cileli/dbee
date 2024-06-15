package main

import (
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var formConnectionNew *tview.Form
var listSavedConnections *tview.List
var flexConnection *tview.Flex
var databaseDrivers = []string{"MySQL", "PostgreSQL"}

type pageConnectionType struct{}

func (pageConnection *pageConnectionType) build() *pageConnectionType {
	listSavedConnections = tview.NewList()
	listSavedConnections.SetBorder(true).SetTitle("Saved Connections (alt+s)").SetTitleAlign(tview.AlignCenter)
	listSavedConnections.SetBorderPadding(1, 1, 2, 2)

	savedConnections := application.getSavedConnections()
	if len(savedConnections) > 0 {
		listShortcutIndex := 1
		shortcutRune := '1'
		for _, savedConnection := range savedConnections {
			connectionString := strings.Split(savedConnection, "|||")
			if len(connectionString) != 2 {
				continue
			}
			listSavedConnections.AddItem(connectionString[1], connectionString[0], rune(shortcutRune), nil)
			if listShortcutIndex < len(listShortcuts) {
				shortcutRune = listShortcuts[listShortcutIndex]
				listShortcutIndex++
			}
		}
	}

	formConnectionNew = tview.NewForm().
		AddDropDown("Driver", databaseDrivers, 0, nil).
		AddInputField("Connection (*)", "", 0, nil, nil).
		AddPasswordField("Password", "", 0, '*', nil).
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

	formConnectionNew.SetBorder(true).SetTitle("Connect to Server (alt+d)").SetTitleAlign(tview.AlignCenter)
	formConnectionNew.SetButtonsAlign(tview.AlignCenter)
	formConnectionNew.SetBorderPadding(1, 1, 2, 2)

	flexConnection = tview.NewFlex().
		AddItem(listSavedConnections, 0, 3, true).
		AddItem(formConnectionNew, 0, 5, true)

	flexConnection.SetBorderPadding(1, 1, 2, 2)

	listSavedConnections.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			if listSavedConnections.GetCurrentItem() == -1 {
				return event
			}
			mainText, secondaryText := listSavedConnections.GetItemText(listSavedConnections.GetCurrentItem())
			formConnectionNew.GetFormItemByLabel("Driver").(*tview.DropDown).SetCurrentOption(slices.Index(databaseDrivers, secondaryText))
			formConnectionNew.GetFormItemByLabel("Connection (*)").(*tview.InputField).SetText(mainText)
			formConnectionNew.GetFormItemByLabel("Password").(*tview.InputField).SetText("")
			app.SetFocus(formConnectionNew)
			app.SetFocus(formConnectionNew.GetFormItemByLabel("Password"))
		}
		return event
	})

	flexConnection.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 's' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(listSavedConnections)
		}
		if event.Rune() == 'd' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(formConnectionNew)
		}
		return event
	})

	if listSavedConnections.GetItemCount() > 0 {
		app.SetFocus(listSavedConnections)
		listSavedConnections.SetCurrentItem(0)
	} else {
		app.SetFocus(formConnectionNew)
		app.SetFocus(formConnectionNew.GetFormItemByLabel("Connection (*)"))
		formConnectionNew.GetFormItemByLabel("Connection (*)").(*tview.InputField).SetText("USER@tcp(HOST:PORT)/DBNAME")
	}

	pages.AddPage("connection", flexConnection, true, false)

	return pageConnection
}

func (pageConnection *pageConnectionType) show() {
	pages.ShowPage("connection")
}
