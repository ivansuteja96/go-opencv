package cpu

/*
#include <stdlib.h>
#include "core.hpp"
*/
import "C"
import (
	"image"
	"unsafe"

	"github.com/ivansuteja96/go-opencv/entity"
)

func NewOpenCVCPU() (method entity.OpenCVMethod) {
	return OpenCV{}
}

func (opencv OpenCV) NewMat() (mat entity.MatMethod) {
	mat = entity.Mat{
		ProcessingUnit: entity.OPENCV_CPU,
		MatMethod: Mat{
			mat: C.NewMat(),
		},
	}
	return
}

//ImageToMatRGB converts image.Image to gocv.Mat (RGB),
//Type of Mat is gocv.MatTypeCV8UC3.
func (opencv OpenCV) ImageToMatRGB(img image.Image) (mat entity.MatMethod, err error) {
	data, width, height := entity.ImageToBytesRGB(img)
	return opencv.NewMatFromBytes(height, width, entity.MatTypeCV8UC3, data)
}

// NewMatFromBytes returns a new Mat with a specific size and type, initialized from a []byte.
func (opencv OpenCV) NewMatFromBytes(rows int, cols int, mt entity.MatType, data []byte) (mat entity.MatMethod, err error) {
	cBytes, err := entity.BytestoCBytes(data)
	if err != nil {
		return Mat{}, err
	}

	mat = entity.Mat{
		ProcessingUnit: entity.OPENCV_CPU,
		MatMethod: Mat{
			mat: C.NewMatFromBytes(C.int(rows), C.int(cols), C.int(mt), *(*C.struct_ByteArray)(unsafe.Pointer(cBytes))),
		},
	}

	return mat, nil
}
