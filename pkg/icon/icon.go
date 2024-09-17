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
	"image/png"
	"os"
	"unsafe"

	"github.com/mat/besticon/ico"
)

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
