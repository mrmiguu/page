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

	date := sock.Rstring()

	for d := range date {
		println(`"` + d + `"`)
	}
}
