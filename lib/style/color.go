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

// func AnsiToRGB(ansi int) []int {
// 	lowRGB := [16][3]int{
// 		{0, 0, 0}, {128, 0, 0}, {0, 128, 0}, {128, 128, 0},
// 		{0, 0, 128}, {128, 0, 128}, {0, 128, 128}, {192, 192, 192},
// 		{128, 128, 128}, {255, 0, 0}, {0, 255, 0}, {255, 255, 0},
// 		{0, 0, 255}, {255, 0, 255}, {0, 255, 255}, {255, 255, 255},
// 	}

// 	if ansi < 0 || ansi > 255 {
// 		return []int{0, 0, 0}
// 	}

// 	if ansi < 16 {
// 		return []int{lowRGB[ansi][0], lowRGB[ansi][1], lowRGB[ansi][2]}
// 	}

// 	if ansi > 231 {
// 		s := (ansi-232)*10 + 8
// 		return []int{s, s, s}
// 	}

// 	n := ansi - 16
// 	b := n % 6
// 	g := (n - b) / 6 % 6
// 	r := (n - b - g*6) / 36 % 6
// 	b = b*40 + 55
// 	r = r*40 + 55
// 	g = g*40 + 55

// 	return []int{r, g, b}
// }
