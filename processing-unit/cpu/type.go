package cpu

/*
#include "core.hpp"
*/
import "C"

type OpenCV struct {
}

type Mat struct {
	mat C.Mat
}
