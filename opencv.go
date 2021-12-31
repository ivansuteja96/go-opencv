package opencv

import (
	"github.com/ivansuteja96/go-opencv/entity"
	"github.com/ivansuteja96/go-opencv/processing-unit/cpu"
	"github.com/ivansuteja96/go-opencv/processing-unit/gpu"
)

func New(cfg Config) (m entity.OpenCVMethod) {
	if cfg.ProcessingUnit == entity.OPENCV_AUTO {
		cfg.ProcessingUnit = entity.OPENCV_CPU
	}

	if cfg.ProcessingUnit == entity.OPENCV_CPU {
		m = cpu.NewOpenCVCPU()
	} else {
		m = gpu.NewOpenCVGPU()
	}
	return
}
