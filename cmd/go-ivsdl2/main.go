package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	formats "github.com/aurelien-cuvelier/go-ivsdl2/internal/formats"
)

var extensionsIsSupported = map[string]bool{
	"bpm": true,
}
var extensionsArr = make([]string, 0, len(extensionsIsSupported))

var loaders = map[string]func(file *os.File) (formats.RawImage, error){
	"bmp": formats.Process,
}

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

	extension := filepath.Ext(imgPath)

	populateExtensionArr()

	if !extensionsIsSupported[extension] {
		fmt.Printf("File format %s is not supported. Supported extensions: %s\n", extension, strings.Join(extensionsArr, ","))
		return
	}

}

func populateExtensionArr() {

	for k := range extensionsIsSupported {
		extensionsArr = append(extensionsArr, k)
	}
}
