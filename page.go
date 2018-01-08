package page

import (
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

func catch(err *error) {
	if e, ok := recover().(*js.Error); ok {
		*err = e
	}
}

type Elem struct {
	js *js.Object
}

func Elems(class string) []Elem {
	elems := []Elem{}
	arr := js.Global.Get("document").Call("getElementsByClassName", class)
	for i := 0; i < arr.Length(); i++ {
		elems = append(elems, Elem{arr.Index(i)})
	}
	return elems
}

func (e *Elem) Translate(x, y int) {
	X, Y := strconv.Itoa(x), strconv.Itoa(y)
	e.js.Get("style").Set("transform", "translate("+X+"px,"+Y+"px)")
}

func (e *Elem) Scale(sx, sy float64) {
	Sx, Sy := strconv.FormatFloat(sx, 'f', 14, 64), strconv.FormatFloat(sy, 'f', 14, 64)
	e.js.Get("style").Set("transform", "scale("+Sx+","+Sy+")")
}

func (e *Elem) Rotate(a float64) {
	A := strconv.FormatFloat(a, 'f', 14, 64)
	e.js.Get("style").Set("transform", "rotate("+A+"deg)")
}

func (e *Elem) Opacity(a float64) {
	A := strconv.FormatFloat(a, 'f', 14, 64)
	e.js.Get("style").Set("opacity", A)
}

// func ID(id string) (e Elem, err error) {
// 	defer catch(&err)
// 	e = Elem{js.Global.Get("document").Call("getElementById", id)}
// 	if e.js == nil {
// 		err = errors.New("bad id")
// 	}
// 	return
// }

// type Audio struct {
// 	js *js.Object
// }

// func GetAudio(id string) (a Audio, err error) {
// 	defer catch(&err)
// 	a = Audio{js.Global.Get("document").Call("getElementById", id)}
// 	if a.js == nil {
// 		err = errors.New("GetAudio: bad id")
// 		return
// 	}
// 	if a.js.Get("tagName").String() != "AUDIO" {
// 		err = errors.New("GetAudio: bad tag")
// 		return
// 	}
// 	return
// }

// func (e *Elem) Anim(class string) (done <-chan bool, err error) {
// 	defer catch(&err)
// 	c := make(chan bool, 1)
// 	e.js.Call("addEventListener", "animationend", func() { c <- true }, false)
// 	println(`before: e.js.Get("className").String()`, e.js.Get("className").String())
// 	e.js.Set("className", e.js.Get("className").String()+" "+class)
// 	println(`after:  e.js.Get("className").String()`, e.js.Get("className").String())
// 	elem := js.Global.Get("document").Call("querySelector", "."+class)
// 	style := js.Global.Get("window").Call("getComputedStyle", elem)
// 	if name := style.Get("animation-name").String(); name == "none" {
// 		err = errors.New(class + " animation not found")
// 		return
// 	}
// 	done = c
// 	return
// }
