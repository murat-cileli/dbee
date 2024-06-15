package main

import (
	"github.com/rivo/tview"
)

var application applicationType
var pages *tview.Pages
var pageConnection pageConnectionType
var pageMain pageMainType
var pageAlert pageAlertType

var database databaseType

func main() {
	application.init()
}
