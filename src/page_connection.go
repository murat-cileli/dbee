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

	formConnectionNew.SetBorder(true).SetTitle("Connect to Server (alt+c)").SetTitleAlign(tview.AlignCenter)
	formConnectionNew.SetButtonsAlign(tview.AlignCenter)

	flexConnection = tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(listSavedConnections, 0, 3, true).
		AddItem(formConnectionNew, 0, 5, true).
		AddItem(nil, 0, 1, false)

	pages.AddPage("connection", flexConnection, true, false)

	listSavedConnections.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		formConnectionNew.GetFormItemByLabel("Driver").(*tview.DropDown).SetCurrentOption(slices.Index(databaseDrivers, secondaryText))
		formConnectionNew.GetFormItemByLabel("Connection (*)").(*tview.InputField).SetText(mainText)
		formConnectionNew.GetFormItemByLabel("Password").(*tview.InputField).SetText("")
	})

	flexConnection.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 's' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(listSavedConnections)
		}
		if event.Rune() == 'c' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(formConnectionNew)
		}
		return event
	})

	if len(savedConnections) > 0 {
		app.SetFocus(listSavedConnections)
	} else {
		app.SetFocus(formConnectionNew)
	}

	return pageConnection
}

func (pageConnection *pageConnectionType) show() {
	pages.ShowPage("connection")
}
