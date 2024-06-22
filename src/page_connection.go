package main

import (
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageConnectionType struct {
	formConnectionNew    *tview.Form
	listSavedConnections *tview.List
	flexConnection       *tview.Flex
	databaseDrivers      []string
}

var pageConnection pageConnectionType

func (pageConnection *pageConnectionType) build() *pageConnectionType {
	pageConnection.databaseDrivers = []string{"MySQL/MariaDB", "PostgreSQL"}
	pageConnection.listSavedConnections = tview.NewList()
	pageConnection.listSavedConnections.SetBorder(true).SetTitle("Saved Connections (alt+s)").SetTitleAlign(tview.AlignCenter)
	pageConnection.listSavedConnections.SetBorderPadding(1, 1, 2, 2)

	savedConnections := application.getSavedConnections()
	if len(savedConnections) > 0 {
		listShortcutIndex := 1
		shortcutRune := '1'
		for _, savedConnection := range savedConnections {
			connectionString := strings.Split(savedConnection, " ")
			if len(connectionString) != 4 {
				continue
			}
			pageConnection.listSavedConnections.AddItem(connectionString[2]+"@"+connectionString[1]+"/"+connectionString[3], connectionString[0], rune(shortcutRune), nil)
			if listShortcutIndex < len(application.ListShortcuts) {
				shortcutRune = application.ListShortcuts[listShortcutIndex]
				listShortcutIndex++
			}
		}
	}

	pageConnection.formConnectionNew = tview.NewForm().
		AddDropDown("Driver", pageConnection.databaseDrivers, 0, nil).
		AddInputField("Host (*)", "", 0, nil, nil).
		AddInputField("User (*)", "", 0, nil, nil).
		AddPasswordField("Password", "", 0, '*', nil).
		AddInputField("Database (*)", "", 0, nil, nil).
		AddButton("Connect", func() {
			_, database.DriverName = pageConnection.formConnectionNew.GetFormItemByLabel("Driver").(*tview.DropDown).GetCurrentOption()
			database.Host = pageConnection.formConnectionNew.GetFormItemByLabel("Host (*)").(*tview.InputField).GetText()
			database.User = pageConnection.formConnectionNew.GetFormItemByLabel("User (*)").(*tview.InputField).GetText()
			database.Password = pageConnection.formConnectionNew.GetFormItemByLabel("Password").(*tview.InputField).GetText()
			database.Database = pageConnection.formConnectionNew.GetFormItemByLabel("Database (*)").(*tview.InputField).GetText()
			err := database.Connect()
			if err != nil {
				pageAlert.show(err.Error(), "error")
				return
			} else {
				if pageConnection.formConnectionNew.GetFormItemByLabel("Save connection").(*tview.Checkbox).IsChecked() {
					application.saveConnection()
				}
				pageMain.build()
				application.pages.ShowPage("main")
				pageMain.focusQueryBox()
			}
		}).
		AddButton("Quit", func() {
			application.ConfirmQuit()
		}).
		AddCheckbox("Save connection", false, nil)

	pageConnection.formConnectionNew.SetBorder(true).SetTitle("Connect to Server (alt+d)").SetTitleAlign(tview.AlignCenter)
	pageConnection.formConnectionNew.SetButtonsAlign(tview.AlignCenter)
	pageConnection.formConnectionNew.SetBorderPadding(1, 1, 2, 2)

	pageConnection.flexConnection = tview.NewFlex().
		AddItem(pageConnection.listSavedConnections, 0, 3, true).
		AddItem(pageConnection.formConnectionNew, 0, 5, true)

	pageConnection.flexConnection.SetBorderPadding(1, 1, 2, 2)

	pageConnection.listSavedConnections.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if pageConnection.listSavedConnections.GetItemCount() == 0 {
			return event
		}
		if event.Key() == tcell.KeyEnter {
			// user@host/database
			mainText, secondaryText := pageConnection.listSavedConnections.GetItemText(pageConnection.listSavedConnections.GetCurrentItem())
			mainTextPartsUserAndHost := strings.Split(mainText, "@")
			mainTextPartsHostAndDatabase := strings.Split(mainTextPartsUserAndHost[1], "/")
			pageConnection.formConnectionNew.GetFormItemByLabel("Driver").(*tview.DropDown).SetCurrentOption(slices.Index(pageConnection.databaseDrivers, secondaryText))
			pageConnection.formConnectionNew.GetFormItemByLabel("Host (*)").(*tview.InputField).SetText(mainTextPartsHostAndDatabase[0])
			pageConnection.formConnectionNew.GetFormItemByLabel("User (*)").(*tview.InputField).SetText(mainTextPartsUserAndHost[0])
			pageConnection.formConnectionNew.GetFormItemByLabel("Password").(*tview.InputField).SetText("")
			pageConnection.formConnectionNew.GetFormItemByLabel("Database (*)").(*tview.InputField).SetText(mainTextPartsHostAndDatabase[1])
			app.SetFocus(pageConnection.formConnectionNew)
			app.SetFocus(pageConnection.formConnectionNew.GetFormItemByLabel("Password"))
		}
		return event
	})

	pageConnection.flexConnection.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 's' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(pageConnection.listSavedConnections)
		}
		if event.Rune() == 'd' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(pageConnection.formConnectionNew)
		}
		return event
	})

	if pageConnection.listSavedConnections.GetItemCount() > 0 {
		app.SetFocus(pageConnection.listSavedConnections)
		pageConnection.listSavedConnections.SetCurrentItem(0)
	} else {
		app.SetFocus(pageConnection.formConnectionNew)
		app.SetFocus(pageConnection.formConnectionNew.GetFormItemByLabel("Connection (*)"))
	}

	application.pages.AddPage("connection", pageConnection.flexConnection, true, false)

	return pageConnection
}

func (pageConnection *pageConnectionType) show() {
	application.pages.ShowPage("connection")
}
