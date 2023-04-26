package utils

import (
	"ascii-art/lib/style"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var banner = new(string)
var Text = new(string)
var asciiArt = new([][8]string)

func ReadASCIIArtFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var art [][8]string
	var currentChar [8]string
	var n int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			art = append(art, currentChar)
			currentChar = [8]string{}
			n = 0
		} else {
			currentChar[n] = line
			n++
		}
	}
	art = append(art, currentChar)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	*asciiArt = art
}

func GetCharIndex(r rune) int {
	if r < 32 || r > 126 {
		fmt.Printf("\033[0G%vError:\033[0m Character [%v%v\033[0m] not available in ascii-art.\n", style.GetColorByName("red"), style.GetColorByName("yellow"), string(r))
		os.Exit(1)
	}
	b := int(r) - 31
	return b
}

func PrintWordAsciiArt() string {
	*banner = "standard"
	output := ""

	path := fmt.Sprintf("data/%v.txt", *banner)

	input := strings.ReplaceAll(*Text, "\\n", "\n")
	lines := strings.Split(input, "\n")

	ReadASCIIArtFile(path)

	for xw, line := range lines {
		if line == "" {
			if xw < len(lines)-1 {
				output += "\n"
				fmt.Println()
			}
		} else {
			words := strings.Split(line, " ")

			for i := 0; i < 8; i++ {

				for x, word := range words {

					for _, v := range word {
	
						idx := GetCharIndex(v)

						asciiChar := (*asciiArt)[idx]

						output += asciiChar[i]
						fmt.Printf("%v", asciiChar[i])

					}
					if  x < len(words)-1 {
						temp := strings.Repeat(" ", 6)
						output += temp
						fmt.Print(strings.Repeat(" ", 6))
					} 

				}
				if i < 8 {
					output += "\n"
					fmt.Println()
				}
			}
		}
	}
	return output
}

func IsIn(let []rune, r rune) bool {
	for _, v := range let {
		if v == r {
			return true
		}
	}
	return false
}

/*

	func PrintUsage() {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	}

	func isPrintableChar(r rune) bool {
		if r >= 32 && r <= 126 {
			return true
		}
		return false
	}
*/

