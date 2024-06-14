package main

import (
	"database/sql"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var pages *tview.Pages
var loginContainer *tview.Flex
var loginForm *tview.Form
var mainContainer *tview.Flex
var alertError *tview.Modal
var textAreaQuery *tview.TextArea
var tableResult *tview.Table
var database databaseType

func main() {
	app = tview.NewApplication()
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorNone.TrueColor()

	buildLoginPage()
	buildAlertErrorPage()

	pages.ShowPage("login")

	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

func buildLoginPage() {
	loginForm = tview.NewForm().
		AddDropDown("Driver", []string{"MariaDB / MySQL"}, 0, nil).
		AddInputField("Host (*)", "localhost", 0, nil, nil).
		AddInputField("Port (*)", "3306", 10, nil, nil).
		AddInputField("User (*)", "root", 0, nil, nil).
		AddPasswordField("Password", "root", 0, '*', nil).
		AddInputField("Database", "mydb", 0, nil, nil).
		AddButton("Connect", func() {
			database.Host = loginForm.GetFormItemByLabel("Host (*)").(*tview.InputField).GetText()
			database.Port = loginForm.GetFormItemByLabel("Port (*)").(*tview.InputField).GetText()
			database.User = loginForm.GetFormItemByLabel("User (*)").(*tview.InputField).GetText()
			database.Password = loginForm.GetFormItemByLabel("Password").(*tview.InputField).GetText()
			database.Database = loginForm.GetFormItemByLabel("Database").(*tview.InputField).GetText()
			err := database.Connect()
			if err != nil {
				alertError.SetText(err.Error())
				alertError.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					pages.HidePage("alertError")
					pages.ShowPage("login")
				})
				pages.ShowPage("alertError")
				return
			} else {
				buildMainPage()
				pages.ShowPage("main")
				app.SetFocus(textAreaQuery)
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		}).
		AddCheckbox("Save connection", false, nil)

	loginForm.SetBorder(true).SetTitle("Connect to Server").SetTitleAlign(tview.AlignCenter)
	loginForm.SetButtonsAlign(tview.AlignCenter)

	loginContainer = tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(loginForm, 0, 2, true).
		AddItem(nil, 0, 1, false)

	pages = tview.NewPages().AddPage("login", loginContainer, true, false)
}

func buildMainPage() {
	listObjects := tview.NewList()

	listObjects.
		SetBorder(true).
		SetTitle("Objects (alt+o)").
		SetTitleAlign(tview.AlignCenter)

	loadObjects(listObjects)

	textAreaQuery = tview.NewTextArea()

	textAreaQuery.
		SetPlaceholder("Type your query here, (alt + return) to run.").
		SetBorder(true).
		SetTitle("Query (alt+q)").
		SetTitleAlign(tview.AlignCenter)

	textAreaQuery.SetText("SELECT * FROM table1", true)

	tableResult = tview.NewTable()
	tableResult.SetTitle("Result")
	tableResult.SetBorders(true)

	mainContainer = tview.NewFlex().
		AddItem(listObjects, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(textAreaQuery, 0, 1, true).
			AddItem(tableResult, 0, 4, true).
			AddItem(tview.NewBox(), 0, 0, false), 0, 2, false).
		AddItem(tview.NewBox(), 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter && event.Modifiers() == tcell.ModAlt {
			event = nil
			result := database.Query(textAreaQuery.GetText())
			loadResult(result)
		}
		return event
	})

	pages.AddPage("main", mainContainer, true, false)
}

func buildAlertErrorPage() {
	alertError = tview.NewModal().
		AddButtons([]string{"OK"})
	alertError.Box.SetBackgroundColor(tcell.ColorDarkRed)
	alertError.SetBackgroundColor(tcell.ColorDarkRed)
	alertError.SetBorder(true).SetTitle("Error").SetTitleAlign(tview.AlignCenter)
	pages.AddPage("alertError", alertError, true, false)
}

func loadObjects(listObjects *tview.List) {
	objectsShortcuts := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}
	objectShortcutIndex := 0
	shortcutRune := rune(0)

	tables := database.getTables()
	for tables.Next() {
		table := ""
		tables.Scan(&table)
		shortcutRune = rune(0)
		if objectShortcutIndex < len(objectsShortcuts) {
			shortcutRune = objectsShortcuts[objectShortcutIndex]
			objectShortcutIndex++
		}
		listObjects.
			AddItem(table, "", shortcutRune, nil).
			ShowSecondaryText(false)
	}
}

func loadResult(rows *sql.Rows) {
	tableResult.Clear()

	if rows == nil {
		return
	}

	columns, _ := rows.Columns()
	for i, column := range columns {
		tableResult.SetCellSimple(0, i, column)
	}

	columnsCount := len(columns)
	values := make([]sql.NullString, columnsCount)
	valuePtrs := make([]any, columnsCount)
	rowCount := 1

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			panic(err)
		}

		for i, cell := range values {
			tableResult.SetCellSimple(rowCount, i, cell.String)
		}

		rowCount++
	}

	tableResult.SetFixed(1, 0)

}
