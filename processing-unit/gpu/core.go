package gpu

import (
	"image"

	"github.com/ivansuteja96/go-opencv/entity"
)

func NewOpenCVGPU() (method entity.OpenCVMethod) {
	return OpenCV{}
}

func (opencv OpenCV) NewMat() (mat entity.MatMethod) {
	mat = entity.Mat{
		ProcessingUnit: entity.OPENCV_GPU,
		MatMethod:      Mat{
			// mat: C.NewMat(),
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
	// cBytes, err := toByteArray(data)
	// if err != nil {
	// 	return Mat{}, err
	// }
	mat = entity.Mat{
		ProcessingUnit: entity.OPENCV_GPU,
		MatMethod:      Mat{
			// mat: C.NewMatFromBytes(C.int(rows), C.int(cols), C.int(mt), *cBytes),
		},
	}

	return mat, nil
}
