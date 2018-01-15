package css

import (
	"strconv"
)

type Display string
type Position string
type Length string

const (
	None Display = "none"
	Grid Display = "grid"

	Absolute Position = "absolute"
	Relative Position = "relative"
	Fixed    Position = "fixed"

	Z Length = "0"
)

func itol(i int, end string) Length {
	return Length(strconv.Itoa(i) + end)
}
func ftol(n float64, end string) Length {
	return Length(strconv.FormatFloat(n, 'f', 14, 64) + end)
}

func Pct(n float64) Length {
	return ftol(n, "%")
}

func Ch(n float64) Length {
	return ftol(n, "ch")
}
func Em(n float64) Length {
	return ftol(n, "em")
}
func Ex(n float64) Length {
	return ftol(n, "ex")
}
func Rem(n float64) Length {
	return ftol(n, "rem")
}

func Vh(n float64) Length {
	return ftol(n, "vh")
}
func Vw(n float64) Length {
	return ftol(n, "vw")
}
func Vmin(n float64) Length {
	return ftol(n, "vmin")
}
func Vmax(n float64) Length {
	return ftol(n, "vmax")
}

func Px(i int) Length {
	return itol(i, "px")
}
func Mm(n float64) Length {
	return ftol(n, "mm")
}
func Cm(n float64) Length {
	return ftol(n, "cm")
}
func In(n float64) Length {
	return ftol(n, "in")
}
func Pt(n float64) Length {
	return ftol(n, "pt")
}
func Pc(n float64) Length {
	return ftol(n, "pc")
}

func Deg(n float64) Length {
	return ftol(n, "deg")
}
