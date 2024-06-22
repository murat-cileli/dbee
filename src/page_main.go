package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageMainType struct {
	flexQuerySize       int
	flexMain            *tview.Flex
	flexInner           *tview.Flex
	textAreaQuery       *tview.TextArea
	listDatabaseObjects *tview.List
	pages               *tview.Pages
}

var pageMain pageMainType

func (pageMain *pageMainType) build() {
	pageMain.listDatabaseObjects = tview.NewList()

	pageMain.listDatabaseObjects.
		SetBorder(true).
		SetTitle("Objects (alt+w)").
		SetTitleAlign(tview.AlignCenter)

	pageMain.listDatabaseObjects.SetSelectedFocusOnly(true)

	pageMain.loadDatabaseObjects()

	pageMain.textAreaQuery = tview.NewTextArea()

	pageMain.textAreaQuery.
		SetPlaceholder("Type your query here, (alt+enter) to execute.").
		SetBorder(true).
		SetTitle(" [yellow]" + database.Database + "[-:-:-:-] |" + " Query (alt+e) ").
		SetTitleAlign(tview.AlignCenter)

	pageMain.pages = tview.NewPages()
	pageMainTable.build()
	pageMainMessage.build()
	pageMainMessage.show(tview.AlignCenter, "", "helpText")

	pageMain.flexQuerySize = 1

	pageMain.flexInner = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(pageMain.textAreaQuery, 0, 1, true).
		AddItem(pageMain.pages, 0, 4, true)

	pageMain.flexMain = tview.NewFlex().
		AddItem(pageMain.listDatabaseObjects, 0, 1, false).
		AddItem(pageMain.flexInner, 0, 2, false)

	pageMain.flexMain.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'w' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(pageMain.listDatabaseObjects)
			return nil
		}
		if event.Rune() == 'e' && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(pageMain.textAreaQuery)
		}
		if event.Rune() == 'r' && event.Modifiers() == tcell.ModAlt {
			pageMainTable.focusTable()
			pageMainTable.tableQueryResults.SetSelectable(true, false)
		}
		if event.Rune() == 'h' && event.Modifiers() == tcell.ModAlt {
			pageMain.pages.SwitchToPage("message")
			pageMainMessage.focus()
			pageMainMessage.show(tview.AlignLeft, "", "helpText")
		}
		if event.Rune() == 'm' && event.Modifiers() == tcell.ModAlt {
			pageMain.resizeFlexQuery("down")
		}
		if event.Rune() == 'j' && event.Modifiers() == tcell.ModAlt {
			pageMain.resizeFlexQuery("up")
		}
		return event
	})

	pageMain.textAreaQuery.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter && event.Modifiers() == tcell.ModAlt {
			event = nil
			pageMain.textAreaQuery.SetDisabled(true)
			results, err := database.Query(pageMain.textAreaQuery.GetText(), true)
			if err == nil {
				pageMainTable.loadQueryResults(results)
			}
			pageMain.textAreaQuery.SetDisabled(false)
		} else if event.Key() == tcell.KeyUp && event.Modifiers() == tcell.ModAlt {
			queryHistory := queryHistory.back()
			if queryHistory != "" {
				pageMain.textAreaQuery.SetText(queryHistory, true)
			}
		} else if event.Key() == tcell.KeyDown && event.Modifiers() == tcell.ModAlt {
			queryHistory := queryHistory.forward()
			if queryHistory != "" {
				pageMain.textAreaQuery.SetText(queryHistory, true)
			}
		}
		return event
	})

	pageMain.listDatabaseObjects.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			pageMain.browseDatabaseObject()
		}
		if event.Key() == tcell.KeyCtrlSpace {
			pageMain.describeDatabaseObject()
		}
		return event
	})

	application.pages.AddPage("main", pageMain.flexMain, true, false)
}

func (pageMain *pageMainType) resizeFlexQuery(direction string) {
	if direction == "down" && pageMain.flexQuerySize < 4 {
		pageMain.flexQuerySize++
	}
	if direction == "up" && pageMain.flexQuerySize > 1 {
		pageMain.flexQuerySize--
	}
	pageMain.flexInner.ResizeItem(pageMain.textAreaQuery, 0, pageMain.flexQuerySize)
	app.Sync().ForceDraw()

}

func (pageMain *pageMainType) describeDatabaseObject() {
	selectedObject, _ := pageMain.listDatabaseObjects.GetItemText(pageMain.listDatabaseObjects.GetCurrentItem())
	query := ""
	if database.DriverName == "MySQL/MariaDB" {
		query = "DESCRIBE " + selectedObject
	} else if database.DriverName == "PostgreSQL" {
		query = "SELECT * FROM information_schema.columns WHERE table_name = '" + selectedObject + "'"
	}
	results, err := database.Query(query, false)
	if err == nil {
		pageMainTable.loadQueryResults(results)
	}
}

func (pageMain *pageMainType) browseDatabaseObject() {
	selectedObject, _ := pageMain.listDatabaseObjects.GetItemText(pageMain.listDatabaseObjects.GetCurrentItem())
	results, err := database.Query("SELECT * FROM "+selectedObject+" LIMIT 5", false)
	if err == nil {
		pageMainTable.loadQueryResults(results)
	}
}

func (pageMain *pageMainType) loadDatabaseObjects() {
	pageMain.listDatabaseObjects.Clear()

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
		if listShortcutIndex < len(application.ListShortcuts) {
			shortcutRune = application.ListShortcuts[listShortcutIndex]
			listShortcutIndex++
		}
		pageMain.listDatabaseObjects.
			AddItem(table, "", shortcutRune, nil).
			ShowSecondaryText(false)
	}
}

func (pageMain *pageMainType) focusQueryBox() {
	app.SetFocus(pageMain.textAreaQuery)
}
