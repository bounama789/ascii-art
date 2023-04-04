package utils

import (
	"ascii-art/lib/style"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var color = new(string)
var banner = new(string)
var align = new(string)
var output = new(string)
var reverse = new(string)
var text = new(string)
var letters = new(string)

func ReadASCIIArtFile(filePath string) [][]string {
	a,_ := os.Getwd()
	println(a)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var asciiArt [][]string
	var currentChar []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			asciiArt = append(asciiArt, currentChar)
			currentChar = []string{}
		} else {
			currentChar = append(currentChar, line)
		}
	}
	asciiArt = append(asciiArt, currentChar)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return asciiArt
}

func GetCharIndex(r rune) int {
	b := int(r) - 32
	return b
}

func PrintWordAsciiArt() {
	path := fmt.Sprintf("data/%v.txt", *banner)
	color, _ := CheckColorArg()
	outpath := *output
	let := []rune(*letters)
	input := strings.ReplaceAll(*text, "\\n", "\n")
	words := strings.Split(input, "\n")

	asciiArt := ReadASCIIArtFile(path)
	reset := "\033[0m"
	var outfile *os.File

	if outpath != "" {
		temp,_ := os.Create(outpath)
		outfile = temp
	}
	for xw, word := range words {

		for i := 0; i < 8; i++ {

			for _, v := range word {
				idx := GetCharIndex(v)

				asciiChar := asciiArt[idx+1]
				if len(let) > 0 {
					if outpath != "" {
						if IsIn(let, v) {
							fmt.Fprintf(outfile,"%v%v%v ", color, asciiChar[i], reset)
						} else {
							fmt.Fprintf(outfile,"%v ", asciiChar[i])
	
						}
					} else {
					if IsIn(let, v) {
						fmt.Printf("%v%v%v ", color, asciiChar[i], reset)
					} else {
						fmt.Printf("%v ", asciiChar[i])

					}}
				} else {
					if outpath != "" {
						fmt.Fprintf(outfile,"%v%v%v ", color, asciiChar[i], reset)

					} else {
						fmt.Printf("%v%v%v ", color, asciiChar[i], reset)
					}
				}
			}

			if outpath != "" {
				fmt.Fprintln(outfile)

			} else {
				fmt.Println()
			}		
		}
		if outpath != "" {
			fmt.Fprintln(outfile)

		} else if xw < len(words)-1{
			fmt.Println()
		}
	}
}

func CheckColorArg() (string, bool) {
	colorArg := *color
	if len(colorArg) < 1 {
		return "", false
	}
	color := style.GetColorByName(colorArg)
	if color != "" {
		return color,true
	}

	var r, g, b int
	if strings.HasPrefix(colorArg, "#") {
		t1, _ := strconv.ParseInt(colorArg[1:3], 16, 0)
		r = int(t1)
		t2, _ := strconv.ParseInt(colorArg[3:5], 16, 0)
		g = int(t2)
		t3, _ := strconv.ParseInt(colorArg[5:7], 16, 0)
		b = int(t3)
	} else if strings.HasPrefix(colorArg, "rgb(") && strings.HasSuffix(colorArg, ")") {
		parts := strings.Split(colorArg[4:len(colorArg)-1], ",")
		if len(parts) != 3 {
			return "", false
		}
		var err error
		r, err = strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return "", false
		}
		g, err = strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return "", false
		}
		b, err = strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			return "", false
		}
	} else if strings.HasPrefix(colorArg, "hsl(") && strings.HasSuffix(colorArg, ")") {
		parts := strings.Split(colorArg[4:len(colorArg)-1], ",")
		if len(parts) != 3 {
			return "", false
		}
		h, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return "", false
		}
		s, err := strconv.Atoi(strings.TrimSuffix(strings.TrimSpace(parts[1]), "%"))
		if err != nil {
			return "", false
		}
		l, err := strconv.Atoi(strings.TrimSuffix(strings.TrimSpace(parts[2]), "%"))
		if err != nil {
			return "", false
		}
		r, g, b = style.HslToRgb(h, s, l)

	} else {
		code, err := strconv.Atoi(colorArg)
		if err != nil || code < 0 || code > 255 {
			return "", false
		}
		return strconv.Itoa(code), true
	}
	return style.RGBToANSI(r, g, b), true

}

func IsIn(let []rune, r rune) bool {
	for _, v := range let {
		if v == r {
			return true
		}
	}
	return false
}

func GetFlagValue(args []string) (string, string, string, string, string, string) {
	flags := []string{"--color", "--align", "--reverse", "--output"}
	var nflag int

	for i, arg := range args {
		if strings.Contains(arg, "=") {
			flag, value := strings.Split(arg, "=")[0], strings.Split(arg, "=")[1]

			switch flag {
			case flags[0]:
				if len(args)-i >2 &&!strings.Contains(args[i+1],"--") {
					*letters = args[i+1]
					nflag++
				}
				*color = value
				nflag++
			case flags[1]:
				*align = value
				nflag++
			case flags[2]:
				*reverse = value
				nflag++
			case flags[3]:
				*output = value
				nflag++
			}
		} else {
			var value string
			if i < len(args)-1 {
				value = args[i+1]
			}
			switch arg {
			case flags[0]:
				if len(args)-i >3 && !strings.Contains(args[i+2],"--") {
					*letters = args[i+2]
					nflag++
				}
				*color = value
				nflag += 2
			case flags[1]:
				*align = value
				nflag += 2
			case flags[2]:
				*reverse = value
				nflag += 2
			case flags[3]:
				*output = value
				nflag += 2
			}
		}
	}

	temp := args[nflag:]
	*text = temp[0]

	if len(temp) == 2 {
		if IsBanner(temp[1]) {
			*banner = temp[1]
		}
	}
	println(*banner)

	return *color, *align, *reverse, *output, *text, *banner

}

func IsBanner(ban string) bool {
	banners := []string{"standard", "thinkertoy", "shadow"}
	for _, v := range banners {
		if ban == v {
			return true
		}
	}
	return false
}
