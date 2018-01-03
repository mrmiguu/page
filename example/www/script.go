package main

import (
	"github.com/mrmiguu/page"
	"github.com/mrmiguu/sock"
)

func main() {
	parent, err := page.ID("parent")
	if err != nil {
		panic(err)
	}

	parent.Show(true)

	handshake := sock.Wbool()
	handshake <- true
}
