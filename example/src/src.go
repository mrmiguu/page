package src

import (
	"github.com/mrmiguu/jsutil"
	"github.com/mrmiguu/sock"
)

func init() {
	err := jsutil.CompileWithGzip("www/script.go")
	if err != nil {
		panic(err)
	}

	handshakes := sock.Rbool()

	for range handshakes {
		println("a handshake has been made")
	}
}
