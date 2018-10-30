package layers 

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

type Gauss struct {
	Enabled int
	Kernel image.Point
	Sigma float64
}

type Median struct {
	Enabled int
	Kernel int
}

type NoiseSuppression struct {
	Gauss Gauss
	Median Median
}

func NewNoiseSuppression() NoiseSuppression {
	return NoiseSuppression {
		Gauss {
			Enabled: 0,
		},
		Median {
			Enabled: 0,
		},
	}
}

func (ns *NoiseSuppression) AddGauss(kernel image.Point, sigma float64) {
	(*ns).Gauss.Enabled = 1
	(*ns).Gauss.Kernel = kernel
	(*ns).Gauss.Sigma = sigma
}

func (ns *NoiseSuppression) AddMedian(kernel int) {
	(*ns).Median.Enabled = 1
	(*ns).Median.Kernel = kernel
}

func (ns *NoiseSuppression) Supress(mat *gocv.Mat, types []string) {
	for i := 0; i < len(types); i++ {
		if (types[i] == "gauss") {
			if ((*ns).Gauss.Enabled == 0) {
				fmt.Println("Gauss is not enabled")
			}
		} else if (types[i] == "median") {
			if ((*ns).Median.Enabled == 0) {
				fmt.Println("Median is not enabled")
			}
		} else {
			fmt.Println("Unavailable type")
		}
	}
	
	for i := 0; i < len(types); i++ {
		if types[i] == "gauss" {
			gocv.GaussianBlur(*mat, mat, (*ns).Gauss.Kernel, (*ns).Gauss.Sigma, 1.0, 0)
		} else if types[i] == "median" {
			gocv.MedianBlur(*mat, mat, (*ns).Median.Kernel)
		} else {
			fmt.Println("Unavailable type")
		}
	}
}



