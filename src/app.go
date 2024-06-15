package main

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

type applicationType struct{}

var app *tview.Application
var pages *tview.Pages
var pageConnection pageConnectionType
var pageMain pageMainType
var pageAlert pageAlertType
var listShortcuts = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func (application *applicationType) init() {
	app = tview.NewApplication()

	pages = tview.NewPages()
	pageAlert.build()
	pageConnection.build().show()

	if err := app.SetRoot(pages, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

func (application *applicationType) saveConnection() {
	file, err := os.OpenFile(filepath.Join(application.getAppConfigDir(), "connections"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		pageAlert.show(err.Error(), "error")
		return
	}
	defer file.Close()

	if _, err := file.WriteString(database.DriverName + " " + database.Host + " " + database.User + " " + database.Database + "\n"); err != nil {
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

func (application *applicationType) getSavedConnections() []string {
	file, err := os.Open(filepath.Join(application.getAppConfigDir(), "connections"))
	if err != nil {
		return nil
	}
	defer file.Close()

	var connections []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		connections = append(connections, scanner.Text())
	}

	return connections
}
