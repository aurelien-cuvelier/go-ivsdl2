package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No args provided")
		return
	}

	imgPath := os.Args[1]
	fmt.Println("Loading image:", imgPath)

	file, err := os.Open(imgPath)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

}
