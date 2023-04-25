package main

import (
	"ascii-art/lib/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) > 1 {
		fmt.Println("Usage: go run . [STRING] \n\nExample: go run . something")
		return
	}

	*utils.Text = os.Args[1]

	if *utils.Text != "" {
		utils.PrintWordAsciiArt()
	}
}
