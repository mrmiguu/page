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
	if e.js == nil {
		err = errors.New("bad id")
	}
	return
}

type Elem struct {
	js *js.Object
}

// func (e *Elem) Translate(left, top float64) <-chan bool {
// 	c := make(chan bool, 1)
// 	e.js.Call("addEventListener", "transitionend", func() { c <- true }, false)
// 	e.js.Get("style").Set("translate", left, top)
// 	return c
// }
// func (e *Elem) Scale(n float64) <-chan bool {
// 	c := make(chan bool, 1)
// 	e.js.Call("addEventListener", "transitionend", func() { c <- true }, false)
// 	e.js.Get("style").Set("scale", n)
// 	return c
// }
// func (e *Elem) Rotate(r float64) <-chan bool {
// 	c := make(chan bool, 1)
// 	e.js.Call("addEventListener", "transitionend", func() { c <- true }, false)
// 	e.js.Get("style").Set("rotate", r)
// 	return c
// }
func (e *Elem) Anim(class string) (done <-chan bool, err error) {
	defer catch(&err)
	c := make(chan bool, 1)
	e.js.Call("addEventListener", "animationend", func() { c <- true }, false)
	e.js.Set("className", class)
	done = c
	return
}
