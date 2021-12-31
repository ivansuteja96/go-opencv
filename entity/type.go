package entity

import "image"

type (
	OpenCVProcessingUnit int

	OpenCVMethod interface {
		NewMat() (mat MatMethod)
		NewMatFromBytes(rows int, cols int, mt MatType, data []byte) (mat MatMethod, err error)
		ImageToMatRGB(img image.Image) (mat MatMethod, err error)
	}

	MatMethod interface {
	}

	Mat struct {
		ProcessingUnit OpenCVProcessingUnit
		MatMethod      MatMethod
	}

	// MatType is the type for the various different kinds of Mat you can create.
	MatType int
)
