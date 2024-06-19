package main

import (
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageConnectionType struct{}

var pageConnection pageConnectionType
var formConnectionNew *tview.Form
var listSavedConnections *tview.List
var flexConnection *tview.Flex
var databaseDrivers = []string{"MySQL/MariaDB", "PostgreSQL"}

func (pageConnection *pageConnectionType) build() *pageConnectionType {
	listSavedConnections = tview.NewList()
	listSavedConnections.SetBorder(true).SetTitle("Saved Connections (alt+s)").SetTitleAlign(tview.AlignCenter)
	listSavedConnections.SetBorderPadding(1, 1, 2, 2)

	savedConnections := application.getSavedConnections()
	if len(savedConnections) > 0 {
		listShortcutIndex := 1
		shortcutRune := '1'
		for _, savedConnection := range savedConnections {
			connectionString := strings.Split(savedConnection, " ")
			if len(connectionString) != 4 {
				continue
			}
			listSavedConnections.AddItem(connectionString[2]+"@"+connectionString[1]+"/"+connectionString[3], connectionString[0], rune(shortcutRune), nil)
			if listShortcutIndex < len(listShortcuts) {
				shortcutRune = listShortcuts[listShortcutIndex]
				listShortcutIndex++
			}
		}
	}

	formConnectionNew = tview.NewForm().
		AddDropDown("Driver", databaseDrivers, 0, nil).
		AddInputField("Host (*)", "", 0, nil, nil).
		AddInputField("User (*)", "", 0, nil, nil).
		AddPasswordField("Password", "", 0, '*', nil).
		AddInputField("Database (*)", "", 0, nil, nil).
		AddButton("Connect", func() {
			_, database.DriverName = formConnectionNew.GetFormItemByLabel("Driver").(*tview.DropDown).GetCurrentOption()
			database.Host = formConnectionNew.GetFormItemByLabel("Host (*)").(*tview.InputField).GetText()
			database.User = formConnectionNew.GetFormItemByLabel("User (*)").(*tview.InputField).GetText()
			database.Password = formConnectionNew.GetFormItemByLabel("Password").(*tview.InputField).GetText()
			database.Database = formConnectionNew.GetFormItemByLabel("Database (*)").(*tview.InputField).GetText()
			err := database.Connect()
			if err != nil {
				pageAlert.show(err.Error(), "error")
				return
			} else {
				if formConnectionNew.GetFormItemByLabel("Save connection").(*tview.Checkbox).IsChecked() {
					application.saveConnection()
				}
				pageMain.build()
				pages.ShowPage("main")
				app.SetFocus(textAreaQuery)
			}
		}).
		AddButton("Quit", func() {
			application.ConfirmQuit()
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
		if listSavedConnections.GetItemCount() == 0 {
			return event
		}
		if event.Key() == tcell.KeyEnter {
			// user@host/database
			mainText, secondaryText := listSavedConnections.GetItemText(listSavedConnections.GetCurrentItem())
			mainTextPartsUserAndHost := strings.Split(mainText, "@")
			mainTextPartsHostAndDatabase := strings.Split(mainTextPartsUserAndHost[1], "/")
			formConnectionNew.GetFormItemByLabel("Driver").(*tview.DropDown).SetCurrentOption(slices.Index(databaseDrivers, secondaryText))
			formConnectionNew.GetFormItemByLabel("Host (*)").(*tview.InputField).SetText(mainTextPartsHostAndDatabase[0])
			formConnectionNew.GetFormItemByLabel("User (*)").(*tview.InputField).SetText(mainTextPartsUserAndHost[0])
			formConnectionNew.GetFormItemByLabel("Password").(*tview.InputField).SetText("")
			formConnectionNew.GetFormItemByLabel("Database (*)").(*tview.InputField).SetText(mainTextPartsHostAndDatabase[1])
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
	}

	pages.AddPage("connection", flexConnection, true, false)

	return pageConnection
}

func (pageConnection *pageConnectionType) show() {
	pages.ShowPage("connection")
}
