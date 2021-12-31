package entity

/*
#include <stdlib.h>
#include "entity.hpp"
*/
import "C"
import (
	"image"
	"unsafe"
)

//ImageToRGBBytes converts image.Image to []bytes RGB,
//RGB image having 8bit for each component.
func ImageToBytesRGB(img image.Image) (bytes []byte, width int, height int) {
	bounds := img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()
	bytes = make([]byte, 0, width*height*3)
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			bytes = append(bytes, byte(b>>8), byte(g>>8), byte(r>>8))
		}
	}
	return bytes, width, height
}

func BytestoCBytes(b []byte) (*C.struct_ByteArray, error) {
	if len(b) == 0 {
		return nil, ErrEmptyByteSlice
	}
	return &C.struct_ByteArray{
		data:   (*C.char)(unsafe.Pointer(&b[0])),
		length: C.int(len(b)),
	}, nil
}

func IsCudaAvailable() bool {
	return (bool)(C.IsCudaAvailable())
}
