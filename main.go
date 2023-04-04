package main

import (
	"ascii-art/lib/utils"
	"os"
)

func main() {
utils.GetFlagValue(os.Args[1:])
utils.PrintWordAsciiArt()

}
