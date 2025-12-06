package bmp

import (
	"errors"
	"os"
)

type RawImage struct {
	Width  int
	Height int
	Pixels []byte
}

func getRawImage(file *os.File) (img RawImage, err error) {

	return RawImage{}, errors.New("not implemented")

}
