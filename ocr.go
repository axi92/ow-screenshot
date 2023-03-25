package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("screens\\tab\\1_1920x1080_20230324121444.png")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
