package icon

/*
#include <stdlib.h>
#include <stdbool.h>
#include "../../iconExtraction/get-exe-icon.h"
#include "../../iconExtraction/get-exe-icon.c"
*/
import "C"
import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
	"unsafe"

	"github.com/mat/besticon/ico"
)

type Scene struct {
	Width  int
	Height int
	Image  *image.RGBA
}

func NewScene(width, height int) *Scene {
	return &Scene{
		Width:  width,
		Height: height,
		Image:  image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func randomColor() color.RGBA {
	rand := rand.New(rand.NewSource(time.Now().Unix()))

	return color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
}

func (s *Scene) PixelDraw(colorFn func(int, int) color.RGBA) {
	for i := 0; i < s.Width; i++ {
		for j := 0; j < s.Height; j++ {
			s.Image.Set(i, j, colorFn(i, j))
		}
	}
}

func GetIconFromFile(path string, allowEmbeddedPNGs bool) ([]byte, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	var bufLen C.DWORD

	iconData := C.get_exe_icon_from_file_utf8(cPath, C.int(boolToInt(allowEmbeddedPNGs)), &bufLen)

	if iconData == nil {
		return nil, fmt.Errorf("failed to get icon from file")
	}

	iconBytes := C.GoBytes(unsafe.Pointer(iconData), C.int(bufLen))
	C.free(unsafe.Pointer(iconData))

	return iconBytes, nil
}

func DecodeBytesToImage(iconBytes []byte) (image.Image, error) {
	reader := bytes.NewReader(iconBytes)

	return ico.Decode(reader)
}

func SaveAsPNG(filePath string, image image.Image) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, image)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
