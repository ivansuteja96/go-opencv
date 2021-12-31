package entity

import "errors"

const (
	OPENCV_CPU OpenCVProcessingUnit = iota
	OPENCV_GPU
	OPENCV_AUTO
)

const (
	// MatChannels1 is a single channel Mat.
	MatChannels1 = 0

	// MatChannels2 is 2 channel Mat.
	MatChannels2 = 8

	// MatChannels3 is 3 channel Mat.
	MatChannels3 = 16

	// MatChannels4 is 4 channel Mat.
	MatChannels4 = 24
)

const (
	// MatTypeCV8U is a Mat of 8-bit unsigned int
	MatTypeCV8U MatType = 0

	// MatTypeCV8S is a Mat of 8-bit signed int
	MatTypeCV8S = 1

	// MatTypeCV16U is a Mat of 16-bit unsigned int
	MatTypeCV16U = 2

	// MatTypeCV16S is a Mat of 16-bit signed int
	MatTypeCV16S = 3

	// MatTypeCV16SC2 is a Mat of 16-bit signed int with 2 channels
	MatTypeCV16SC2 = MatTypeCV16S + MatChannels2

	// MatTypeCV32S is a Mat of 32-bit signed int
	MatTypeCV32S = 4

	// MatTypeCV32F is a Mat of 32-bit float
	MatTypeCV32F = 5

	// MatTypeCV64F is a Mat of 64-bit float
	MatTypeCV64F = 6

	// MatTypeCV8UC1 is a Mat of 8-bit unsigned int with a single channel
	MatTypeCV8UC1 = MatTypeCV8U + MatChannels1

	// MatTypeCV8UC2 is a Mat of 8-bit unsigned int with 2 channels
	MatTypeCV8UC2 = MatTypeCV8U + MatChannels2

	// MatTypeCV8UC3 is a Mat of 8-bit unsigned int with 3 channels
	MatTypeCV8UC3 = MatTypeCV8U + MatChannels3

	// MatTypeCV8UC4 is a Mat of 8-bit unsigned int with 4 channels
	MatTypeCV8UC4 = MatTypeCV8U + MatChannels4

	// MatTypeCV8SC1 is a Mat of 8-bit signed int with a single channel
	MatTypeCV8SC1 = MatTypeCV8S + MatChannels1

	// MatTypeCV8SC2 is a Mat of 8-bit signed int with 2 channels
	MatTypeCV8SC2 = MatTypeCV8S + MatChannels2

	// MatTypeCV8SC3 is a Mat of 8-bit signed int with 3 channels
	MatTypeCV8SC3 = MatTypeCV8S + MatChannels3

	// MatTypeCV8SC4 is a Mat of 8-bit signed int with 4 channels
	MatTypeCV8SC4 = MatTypeCV8S + MatChannels4

	// MatTypeCV16UC1 is a Mat of 16-bit unsigned int with a single channel
	MatTypeCV16UC1 = MatTypeCV16U + MatChannels1

	// MatTypeCV16UC2 is a Mat of 16-bit unsigned int with 2 channels
	MatTypeCV16UC2 = MatTypeCV16U + MatChannels2

	// MatTypeCV16UC3 is a Mat of 16-bit unsigned int with 3 channels
	MatTypeCV16UC3 = MatTypeCV16U + MatChannels3

	// MatTypeCV16UC4 is a Mat of 16-bit unsigned int with 4 channels
	MatTypeCV16UC4 = MatTypeCV16U + MatChannels4

	// MatTypeCV16SC1 is a Mat of 16-bit signed int with a single channel
	MatTypeCV16SC1 = MatTypeCV16S + MatChannels1

	// MatTypeCV16SC3 is a Mat of 16-bit signed int with 3 channels
	MatTypeCV16SC3 = MatTypeCV16S + MatChannels3

	// MatTypeCV16SC4 is a Mat of 16-bit signed int with 4 channels
	MatTypeCV16SC4 = MatTypeCV16S + MatChannels4

	// MatTypeCV32SC1 is a Mat of 32-bit signed int with a single channel
	MatTypeCV32SC1 = MatTypeCV32S + MatChannels1

	// MatTypeCV32SC2 is a Mat of 32-bit signed int with 2 channels
	MatTypeCV32SC2 = MatTypeCV32S + MatChannels2

	// MatTypeCV32SC3 is a Mat of 32-bit signed int with 3 channels
	MatTypeCV32SC3 = MatTypeCV32S + MatChannels3

	// MatTypeCV32SC4 is a Mat of 32-bit signed int with 4 channels
	MatTypeCV32SC4 = MatTypeCV32S + MatChannels4

	// MatTypeCV32FC1 is a Mat of 32-bit float int with a single channel
	MatTypeCV32FC1 = MatTypeCV32F + MatChannels1

	// MatTypeCV32FC2 is a Mat of 32-bit float int with 2 channels
	MatTypeCV32FC2 = MatTypeCV32F + MatChannels2

	// MatTypeCV32FC3 is a Mat of 32-bit float int with 3 channels
	MatTypeCV32FC3 = MatTypeCV32F + MatChannels3

	// MatTypeCV32FC4 is a Mat of 32-bit float int with 4 channels
	MatTypeCV32FC4 = MatTypeCV32F + MatChannels4

	// MatTypeCV64FC1 is a Mat of 64-bit float int with a single channel
	MatTypeCV64FC1 = MatTypeCV64F + MatChannels1

	// MatTypeCV64FC2 is a Mat of 64-bit float int with 2 channels
	MatTypeCV64FC2 = MatTypeCV64F + MatChannels2

	// MatTypeCV64FC3 is a Mat of 64-bit float int with 3 channels
	MatTypeCV64FC3 = MatTypeCV64F + MatChannels3

	// MatTypeCV64FC4 is a Mat of 64-bit float int with 4 channels
	MatTypeCV64FC4 = MatTypeCV64F + MatChannels4
)

var (
	ErrEmptyByteSlice error = errors.New("empty byte array")
)
