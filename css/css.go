package css

import (
	"strconv"
)

type Display string
type Position string
type Length string
type Time string

const (
	None Display = "none"
	Grid Display = "grid"

	Absolute Position = "absolute"
	Relative Position = "relative"
	Fixed    Position = "fixed"

	Z Length = "0"
)

func itos(i int, end string) string {
	return strconv.Itoa(i) + end
}
func ftos(n float64, end string) string {
	return strconv.FormatFloat(n, 'f', 14, 64) + end
}

func S(n float64) Time {
	return Time(ftos(n, "s"))
}
func Ms(i int) Time {
	return Time(itos(i, "ms"))
}

func Pct(n float64) Length {
	return Length(ftos(n, "%"))
}

func Ch(n float64) Length {
	return Length(ftos(n, "ch"))
}
func Em(n float64) Length {
	return Length(ftos(n, "em"))
}
func Ex(n float64) Length {
	return Length(ftos(n, "ex"))
}
func Rem(n float64) Length {
	return Length(ftos(n, "rem"))
}

func Vh(n float64) Length {
	return Length(ftos(n, "vh"))
}
func Vw(n float64) Length {
	return Length(ftos(n, "vw"))
}
func Vmin(n float64) Length {
	return Length(ftos(n, "vmin"))
}
func Vmax(n float64) Length {
	return Length(ftos(n, "vmax"))
}

func Px(i int) Length {
	return Length(itos(i, "px"))
}
func Mm(n float64) Length {
	return Length(ftos(n, "mm"))
}
func Cm(n float64) Length {
	return Length(ftos(n, "cm"))
}
func In(n float64) Length {
	return Length(ftos(n, "in"))
}
func Pt(n float64) Length {
	return Length(ftos(n, "pt"))
}
func Pc(n float64) Length {
	return Length(ftos(n, "pc"))
}

func Deg(n float64) Length {
	return Length(ftos(n, "deg"))
}
