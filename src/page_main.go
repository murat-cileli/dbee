package main

import (
	"database/sql"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var tableQueryResults *tview.Table
var flexMain *tview.Flex
var textAreaQuery *tview.TextArea
var listDatabaseObjects *tview.List

type pageMainType struct{}

func (pageMain *pageMainType) build() {
	listDatabaseObjects = tview.NewList()

	listDatabaseObjects.
		SetBorder(true).
		SetTitle("Objects (alt+o)").
		SetTitleAlign(tview.AlignCenter)

	pageMain.loadDatabaseObjects()

	textAreaQuery = tview.NewTextArea()

	textAreaQuery.
		SetPlaceholder("Type your query here, (alt + return) to run.").
		SetBorder(true).
		SetTitle("Query (alt+q)").
		SetTitleAlign(tview.AlignCenter)

	textAreaQuery.SetText("SELECT * FROM pootle_store_store", true)

	tableQueryResults = tview.NewTable()
	tableQueryResults.SetTitle("Results (alt+r)")
	tableQueryResults.SetBorders(true)
	tableQueryResults.SetSelectable(true, false)
	tableQueryResults.SetBordersColor(tcell.ColorDimGrey)
	tableQueryResults.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorDarkSlateGrey).Foreground(tcell.ColorWhite))

	flexMain = tview.NewFlex().
		AddItem(listDatabaseObjects, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(textAreaQuery, 0, 1, true).
			AddItem(tableQueryResults, 0, 4, true).
			AddItem(tview.NewBox(), 0, 0, false), 0, 2, false).
		AddItem(tview.NewBox(), 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter && event.Modifiers() == tcell.ModAlt {
			event = nil
			results := database.Query(textAreaQuery.GetText())
			pageMain.loadQueryResults(results)
		}
		return event
	})

	pages.AddPage("main", flexMain, true, false)
}

func (pageMain *pageMainType) loadDatabaseObjects() {
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
		listDatabaseObjects.
			AddItem(table, "", shortcutRune, nil).
			ShowSecondaryText(false)
	}
}

func (pageMain *pageMainType) loadQueryResults(rows *sql.Rows) {
	tableQueryResults.Clear()

	columns, _ := rows.Columns()
	for i, column := range columns {
		tableQueryResults.SetCell(
			0, i,
			&tview.TableCell{
				Text:  column,
				Color: tcell.ColorDarkGoldenrod.TrueColor(),
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
				cellTextColor = tcell.ColorDarkGoldenrod.TrueColor()
			}

			tableQueryResults.SetCell(
				rowCount, i,
				&tview.TableCell{
					Text:  cell.String,
					Color: cellTextColor,
				},
			)
		}

		rowCount++
	}

	if rowCount > 1 && columnsCount > 1 {
		tableQueryResults.SetFixed(1, 1)
	}
}
