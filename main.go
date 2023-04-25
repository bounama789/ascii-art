package main

import (
	"ascii-art/lib/utils"
	"os"
)

func main() {
	if len(os.Args[1:]) > 1 {
	}
	*utils.Text = os.Args[1]

	if *utils.Text != "" {
		utils.PrintWordAsciiArt()
	
	}
}

