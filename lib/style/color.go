package style

import (
	"fmt"
	"math"
)

func RGBToANSI(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func GetColorByName(name string) string {
	colorRgb := COLORS[name]
	if colorRgb != nil {
		return RGBToANSI(colorRgb[0], colorRgb[1], colorRgb[2])

	}
	return ""
}

func HslToRgb(h, s, l int) (int, int, int) {
	var r, g, b float64
	if s == 0 {
		r = float64(l)
		g = float64(l)
		b = float64(l)
	} else {
		s := float64(s) / 100.0
		l := float64(l) / 100.0
		c := (1.0 - math.Abs(2*l-1)) * s
		x := c * (1 - math.Abs(math.Mod(float64(h)/60, 2)-1))
		m := l - c/2
		switch {
		case h < 60:
			r, g, b = c, x, 0
		case h < 120:
			r, g, b = x, c, 0
		case h < 180:
			r, g, b = 0, c, x
		case h < 240:
			r, g, b = 0, x, c
		case h < 300:
			r, g, b = x, 0, c
		case h < 360:
			r, g, b = c, 0, x
		default:
			r, g, b = 0, 0, 0
		}
		r += m
		g += m
		b += m
	}
	return int(r*255 + 0.5), int(g*255 + 0.5), int(b*255 + 0.5)
}
