package main

import (
	"database/sql"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageMainType struct{}

var pageMain pageMainType
var tableQueryResults *tview.Table
var flexMain *tview.Flex
var textAreaQuery *tview.TextArea
var listDatabaseObjects *tview.List

func (pageMain *pageMainType) build() {
	listDatabaseObjects = tview.NewList()

	listDatabaseObjects.
		SetBorder(true).
		SetTitle("Objects (alt+w)").
		SetTitleAlign(tview.AlignCenter)

	listDatabaseObjects.SetSelectedFocusOnly(true)

	pageMain.loadDatabaseObjects()

	textAreaQuery = tview.NewTextArea()

	textAreaQuery.
		SetPlaceholder("Type your query here, (alt+enter) to execute.").
		SetBorder(true).
		SetTitle("Query (alt+q)").
		SetTitleAlign(tview.AlignCenter)

	tableQueryResults = tview.NewTable()
	tableQueryResults.SetTitle("Results (alt+r)")
	tableQueryResults.SetBorders(true)
	tableQueryResults.SetSelectable(false, false)
	tableQueryResults.SetBackgroundColor(tcell.ColorBlack)

	flexMain = tview.NewFlex().
		AddItem(listDatabaseObjects, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(textAreaQuery, 0, 1, true).
			AddItem(tableQueryResults, 0, 4, true).
			AddItem(tview.NewBox(), 0, 0, false), 0, 2, false)

	flexMain.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'w' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(listDatabaseObjects)
		}
		if event.Rune() == 'q' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(textAreaQuery)
		}
		if event.Rune() == 'r' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(tableQueryResults)
			tableQueryResults.SetSelectable(true, false)
		}
		return event
	})

	textAreaQuery.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter && event.Modifiers() == tcell.ModAlt {
			event = nil
			results, err := database.Query(textAreaQuery.GetText())
			if err == nil {
				pageMain.loadQueryResults(results)
			}
		}
		return event
	})

	listDatabaseObjects.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			pageMain.browseDatabaseObject()
		}
		if event.Key() == tcell.KeyCtrlSpace {
			pageMain.describeDatabaseObject()
		}
		return event
	})

	tableQueryResults.SetFocusFunc(func() {
		tableQueryResults.SetSelectable(true, false)
	})

	tableQueryResults.SetBlurFunc(func() {
		tableQueryResults.SetSelectable(false, false)
	})

	pages.AddPage("main", flexMain, true, false)
}

func (pageMain *pageMainType) describeDatabaseObject() {
	selectedObject, _ := listDatabaseObjects.GetItemText(listDatabaseObjects.GetCurrentItem())
	query := ""
	if database.DriverName == "MySQL/MariaDB" {
		query = "DESCRIBE " + selectedObject
	} else if database.DriverName == "PostgreSQL" {
		query = "SELECT * FROM information_schema.columns WHERE table_name = '" + selectedObject + "'"
	}
	results, err := database.Query(query)
	if err == nil {
		pageMain.loadQueryResults(results)
	}
}

func (pageMain *pageMainType) browseDatabaseObject() {
	selectedObject, _ := listDatabaseObjects.GetItemText(listDatabaseObjects.GetCurrentItem())
	results, err := database.Query("SELECT * FROM " + selectedObject + " LIMIT 10")
	if err == nil {
		pageMain.loadQueryResults(results)
	}
}

func (pageMain *pageMainType) loadDatabaseObjects() {
	listDatabaseObjects.Clear()

	tables, err := database.getTables()
	if tables == nil || err != nil {
		return
	}

	listShortcutIndex := 0
	shortcutRune := rune(0)

	for tables.Next() {
		table := ""
		tables.Scan(&table)
		shortcutRune = rune(0)
		if listShortcutIndex < len(listShortcuts) {
			shortcutRune = listShortcuts[listShortcutIndex]
			listShortcutIndex++
		}
		listDatabaseObjects.
			AddItem(table, "", shortcutRune, nil).
			ShowSecondaryText(false)
	}
}

func (pageMain *pageMainType) loadQueryResults(rows *sql.Rows) {
	tableQueryResults.Clear()

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	for i, column := range columns {
		tableQueryResults.SetCell(
			0, i,
			&tview.TableCell{
				Text:            column,
				Color:           tcell.ColorDarkGoldenrod,
				BackgroundColor: tcell.ColorBlack,
			},
		)
	}

	if rows == nil {
		return
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
			cellTextColor := tcell.ColorWhite
			if i == 0 {
				cellTextColor = tcell.ColorDarkGoldenrod
			}

			tableQueryResults.SetCell(
				rowCount, i,
				&tview.TableCell{
					Text:            cell.String,
					Color:           cellTextColor,
					BackgroundColor: tcell.ColorBlack,
				},
			)
		}

		rowCount++
	}

	if rowCount > 1 && columnsCount > 1 {
		tableQueryResults.SetFixed(1, 1)
	}

	tableQueryResults.SetSelectable(false, false)
}
