package main

import (
	"github.com/rivo/tview"
)

type pageMainMessageType struct {
	helpText string
}

var pageMainMessage pageMainMessageType
var textViewMessage *tview.TextView

func (pageMainMessage *pageMainMessageType) build() {
	textViewMessage = tview.NewTextView()
	textViewMessage.SetBorder(true)
	textViewMessage.SetBorderPadding(1, 1, 1, 1)
	textViewMessage.SetTitleAlign(tview.AlignCenter)
	textViewMessage.SetRegions(true)
	textViewMessage.SetWordWrap(true)
	textViewMessage.SetWrap(true)
	textViewMessage.SetScrollable(true)

	pageMainMessage.helpText = `KEYBOARD SHORTCUTS

[green]Global:[-:-:-:-]
[yellow]esc[-:-:-:-]: Quit application.
[yellow]ctrl + shift + v[-:-:-:-]: Paste text.
[yellow]ctrl + z[-:-:-:-]: Undo text.

[green]Main UI:[-:-:-:-]
[yellow]alt + w[-:-:-:-]: Focus on the objects list.
[yellow]alt + q[-:-:-:-]: Focus on the query text area.
[yellow]alt + r[-:-:-:-]: Focus on the results table.
[yellow]alt + h[-:-:-:-]: Show this help message.

[green]Query Pane:[-:-:-:-]
[yellow]alt + enter[-:-:-:-]: Execute the query.
[yellow]alt + up[-:-:-:-]: Go back in the query history.
[yellow]alt + down[-:-:-:-]: Go forward in the query history.

[green]Object List:[-:-:-:-]
[yellow]1..9, a..z[-:-:-:-]: Select an object.
[yellow]enter[-:-:-:-]: Browse top 10 rows of the selected object.
[yellow]ctrl + space[-:-:-:-]: View selected object's structure.

[green]Connections:[-:-:-:-]
[yellow]alt + s[-:-:-:-]: Focus on the saved connections list.
[yellow]1..9, a..z[-:-:-:-]: Select an item from the saved connections list.
[yellow]enter[-:-:-:-]: Apply the selected saved connection.
[yellow]alt + d[-:-:-:-]: Focus on the "Connect to Server" form.
[yellow]tab[-:-:-:-]: Move to the next form field.

ABOUT

DBee is a free and open-source project maintained by contributors. Feel free report issues or submit feature requests on GitHub. Thank you for using DBee!

[yellow]GitHub[-:-:-:-]: [:::https://github.com/murat-cileli/dbee]https://github.com/murat-cileli/dbee[:::-]
`
	pagesMain.AddPage("message", textViewMessage, true, false)
}

func (pageMainMessage *pageMainMessageType) show(textAlign int, title, message string) {
	textViewMessage.Clear()

	switch message {
	case "helpText":
		textViewMessage.SetTextAlign(tview.AlignLeft)
		textViewMessage.SetTitle("DBee Help (alt+h)")
		textViewMessage.SetText(pageMainMessage.helpText)
	default:
		textViewMessage.SetTextAlign(textAlign)
		textViewMessage.SetTitle(title)
		textViewMessage.SetText(message)
	}

	pagesMain.SwitchToPage("message")
}
