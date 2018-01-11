package main

import (
	"github.com/mrmiguu/page"
	"github.com/mrmiguu/sock"
)

var (
	mainScreen = page.Class("main")
	dateInput  = page.Class("date")
	sendButton = page.Class("send")
)

func main() {
	Date := sock.Wstring()

	mainScreen.Display("grid")

	var d string
	for range sendButton.Hit {
		if d = dateInput.Value(); len(d) > 0 {
			break
		}
	}
	dateInput.Disable(true)
	sendButton.Disable(true)

	Date <- d
}
