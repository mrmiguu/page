package page

import (
	"strconv"

	"github.com/mrmiguu/page/css"

	"github.com/gopherjs/gopherjs/js"
)

func crash() {
	if e, ok := recover().(*js.Error); ok {
		panic(e)
	}
}

func catch(err *error) {
	if e, ok := recover().(*js.Error); ok {
		*err = e
	}
}

type Elem struct {
	elems []*js.Object

	Hit,
	Link <-chan bool
}

func callback(len int) (func(), <-chan bool) {
	c := make(chan bool, len)
	return func() {
		go func() {
			select {
			case c <- true:
			default:
			}
		}()
	}, c
}

func elem(elems []*js.Object) Elem {
	hitfn, hit := callback(0)
	link := make(chan bool, 1)

	for _, e := range elems {
		e.Call("addEventListener", "click", hitfn)

		id := e.Get("id").String()
		if id != "#" && len(id) > 0 {
			id = "#" + id
		}
		hash := js.Global.Get("document").Get("location").Get("hash").String()
		if len(hash) == 0 {
			hash = "#"
		}
		if id == hash {
			select {
			case link <- true:
			default:
			}
		}

		js.Global.Get("window").Call("addEventListener", "hashchange", func() {
			go func() {
				hash := js.Global.Get("document").Get("location").Get("hash").String()
				if len(hash) == 0 {
					hash = "#"
				}
				if id == hash {
					select {
					case link <- true:
					default:
					}
				}
			}()
		})
	}

	return Elem{elems, hit, link}
}

func ID(name string) Elem {
	defer crash()
	return elem([]*js.Object{js.Global.Get("document").Call("getElementById", name)})
}

func Class(name string) Elem {
	defer crash()
	arr := js.Global.Get("document").Call("getElementsByClassName", name)
	var js []*js.Object
	for i := 0; i < arr.Length(); i++ {
		js = append(js, arr.Index(i))
	}
	return elem(js)
}

func (e *Elem) Value(s ...string) string {
	if len(s) > 0 {
		e.setvalue(s[0])
	}
	return e.getvalue()
}

func (e *Elem) setvalue(s string) {
	for _, elem := range e.elems {
		elem.Set("value", s)
	}
}

func (e Elem) getvalue() string {
	for _, elem := range e.elems {
		return elem.Get("value").String()
	}
	return ""
}

func (e *Elem) setdisplay(d string) {
	for _, elem := range e.elems {
		elem.Get("style").Set("display", d)
	}
}

func (e Elem) getdisplay() string {
	for _, elem := range e.elems {
		return elem.Get("style").Get("display").String()
	}
	return ""
}

func (e *Elem) Disable(b bool) {
	for _, elem := range e.elems {
		elem.Set("disabled", b)
	}
}

func (e *Elem) Display(d ...css.Display) css.Display {
	if len(d) > 0 {
		e.setdisplay(string(d[0]))
	}
	return css.Display(e.getdisplay())
}

func (e *Elem) Position(p css.Position) {
	defer crash()
	for _, elem := range e.elems {
		elem.Get("style").Set("position", p)
	}
}

func (e *Elem) Animation(name string) {
	defer crash()
	fn, c := callback(0)
	for _, elem := range e.elems {
		elem.Call("addEventListener", "animationend", fn)
		elem.Get("style").Set("animation-name", name)
	}
	<-c
}

func (e *Elem) TranslateY(y css.Length) {
	defer crash()
	fn, c := callback(0)
	for _, elem := range e.elems {
		elem.Call("addEventListener", "transitionend", fn)
		elem.Get("style").Set("transform", "translateY("+y+")")
	}
	<-c
}
func (e *Elem) Translate(x, y css.Length) {
	defer crash()
	fn, c := callback(0)
	for _, elem := range e.elems {
		elem.Call("addEventListener", "transitionend", fn)
		elem.Get("style").Set("transform", "translate("+x+","+y+")")
	}
	<-c
	println("translated.")
}

func (e *Elem) Move(left, top css.Length) {
	defer crash()
	fn, c := callback(0)
	for _, elem := range e.elems {
		elem.Call("addEventListener", "transitionend", fn)
		elem.Get("style").Set("left", left)
		elem.Get("style").Set("top", top)
	}
	<-c
}

func (e *Elem) Scale(sx, sy float64) {
	Sx, Sy := strconv.FormatFloat(sx, 'f', 14, 64), strconv.FormatFloat(sy, 'f', 14, 64)
	for _, elem := range e.elems {
		elem.Get("style").Set("transform", "scale("+Sx+","+Sy+")")
	}
}

func (e *Elem) Rotate(a css.Length) {
	defer crash()
	fn, c := callback(0)
	for _, elem := range e.elems {
		elem.Call("addEventListener", "transitionend", fn)
		elem.Get("style").Set("transform", "rotate("+a+")")
	}
	<-c
}

func (e *Elem) Opacity(a float64) {
	A := strconv.FormatFloat(a, 'f', 14, 64)
	// c := make(chan bool)
	for _, elem := range e.elems {
		// elem.Call("addEventListener", "transitionend", func() { go func() { c <- true }() })
		elem.Get("style").Set("opacity", A)
	}
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
