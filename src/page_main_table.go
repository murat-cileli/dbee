package main

import (
	"database/sql"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageMainTableType struct {
	tableQueryResults *tview.Table
}

var pageMainTable pageMainTableType

func (pageMainTable *pageMainTableType) build() {

	pageMainTable.tableQueryResults = tview.NewTable()
	pageMainTable.tableQueryResults.Box.SetBorder(true)
	pageMainTable.tableQueryResults.Box.SetTitle("Results (alt+r)")
	pageMainTable.tableQueryResults.SetBorders(true)
	pageMainTable.tableQueryResults.SetSelectable(false, false)
	pageMainTable.tableQueryResults.SetBackgroundColor(tcell.ColorBlack)

	pageMainTable.tableQueryResults.SetFocusFunc(func() {
		pageMainTable.tableQueryResults.SetSelectable(true, false)
	})

	pageMainTable.tableQueryResults.SetBlurFunc(func() {
		pageMainTable.tableQueryResults.SetSelectable(false, false)
	})

	pageMain.pages.AddPage("tableQueryResults", pageMainTable.tableQueryResults, true, false)
}

func (pageMainTable *pageMainTableType) loadQueryResults(rows *sql.Rows) {
	pageMainTable.tableQueryResults.Clear()
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	pageMain.pages.SwitchToPage("tableQueryResults")

	for i, column := range columns {
		pageMainTable.tableQueryResults.SetCell(
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

			pageMainTable.tableQueryResults.SetCell(
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
		pageMainTable.tableQueryResults.SetFixed(1, 1)
	}

	pageMainTable.tableQueryResults.SetSelectable(false, false)
}

func (pageMainTable *pageMainTableType) focusTable() {
	app.SetFocus(pageMainTable.tableQueryResults)
}
