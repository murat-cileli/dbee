package main

import (
	"github.com/rivo/tview"
)

type pageMainMessageType struct {
	helpText        string
	textViewMessage *tview.TextView
}

var pageMainMessage pageMainMessageType

func (pageMainMessage *pageMainMessageType) build() {
	pageMainMessage.textViewMessage = tview.NewTextView()
	pageMainMessage.textViewMessage.SetBorder(true)
	pageMainMessage.textViewMessage.SetBorderPadding(1, 1, 1, 1)
	pageMainMessage.textViewMessage.SetTitleAlign(tview.AlignCenter)
	pageMainMessage.textViewMessage.SetRegions(true)
	pageMainMessage.textViewMessage.SetWordWrap(true)
	pageMainMessage.textViewMessage.SetWrap(true)
	pageMainMessage.textViewMessage.SetScrollable(true)

	pageMainMessage.helpText = `KEYBOARD SHORTCUTS

[green]Global:[-:-:-:-]
[yellow]esc[-:-:-:-]: Quit application.
[yellow]ctrl + shift + v[-:-:-:-]: Paste text.
[yellow]ctrl + z[-:-:-:-]: Undo text.

[green]Main UI:[-:-:-:-]
[yellow]alt + w[-:-:-:-]: Focus on the objects list.
[yellow]alt + s[-:-:-:-]: Focus on the query text area.
[yellow]alt + r[-:-:-:-]: Focus on the results table.
[yellow]alt + h[-:-:-:-]: Show this help message.

[green]Query Pane:[-:-:-:-]
[yellow]alt + enter[-:-:-:-]: Execute the query.
[yellow]alt + up[-:-:-:-]: Go back in the query history.
[yellow]alt + down[-:-:-:-]: Go forward in the query history.
[yellow]alt + m[-:-:-:-]: Expand the query text area.
[yellow]alt + j[-:-:-:-]: Shrink the query text area.

[green]Object List:[-:-:-:-]
[yellow]1..9, a..z[-:-:-:-]: Select an object.
[yellow]enter[-:-:-:-]: Browse top 5 rows of the selected object.
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
	pageMain.pages.AddPage("message", pageMainMessage.textViewMessage, true, false)
}

func (pageMainMessage *pageMainMessageType) show(textAlign int, title, message string) {
	pageMainMessage.textViewMessage.Clear()

	switch message {
	case "helpText":
		pageMainMessage.textViewMessage.SetTextAlign(tview.AlignLeft)
		pageMainMessage.textViewMessage.SetTitle("DBee Help (alt+h)")
		pageMainMessage.textViewMessage.SetText(pageMainMessage.helpText)
	default:
		pageMainMessage.textViewMessage.SetTextAlign(textAlign)
		pageMainMessage.textViewMessage.SetTitle(title)
		pageMainMessage.textViewMessage.SetText(message)
	}

	pageMain.pages.SwitchToPage("message")
}

func (pageMainMessage *pageMainMessageType) focus() {
	app.SetFocus(pageMainMessage.textViewMessage)
}
