package main

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type applicationType struct{}

var app *tview.Application

func (application *applicationType) init() {
	app = tview.NewApplication()

	pages = tview.NewPages()
	pageAlert.build()
	pageConnection.build().show()

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorNone.TrueColor()

	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

func (application *applicationType) saveConnection(databaseDriver string, connectionString string) {
	file, err := os.OpenFile(filepath.Join(application.getAppConfigDir(), "connections"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		pageAlert.show(err.Error(), "error")
		return
	}
	defer file.Close()

	if _, err := file.WriteString(databaseDriver + "|||" + connectionString + "\n"); err != nil {
		pageAlert.show(err.Error(), "error")
	}
}

func (application *applicationType) getAppConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	appConfigDir := filepath.Join(configDir, "dbee")
	err = os.MkdirAll(appConfigDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return appConfigDir
}
