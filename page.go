package page

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

func catch(err *error) {
	if e, ok := recover().(*js.Error); ok {
		*err = e
	}
}

func ID(id string) (e Elem, err error) {
	defer catch(&err)
	e = Elem{js.Global.Get("document").Call("getElementById", id)}
	if e.Object == nil {
		err = errors.New("bad id")
	}
	return
}

type Elem struct {
	*js.Object
}

func (e *Elem) Show(b bool) {
	a := 0
	if b {
		a = 1
	}
	e.Get("style").Set("opacity", a)
}
