package layers

import (
//	"image"
	"fmt"
	"gocv.io/x/gocv"
)


type Threshold struct {
	MaxValue float32
	BlockSize int
	C float32
}


func NewThreshold(maxvalue float32, blocksize int, c float32) Threshold {
	return Threshold {
		MaxValue: maxvalue,
		BlockSize: blocksize,
		C: c,
	}
}

func (th *Threshold) Binarize (mat *gocv.Mat, threshtype string) {
	//RGB to GrayScale
	gocv.CvtColor(*mat, mat, 7)
	
	//Equalize Histogramm
	gocv.EqualizeHist(*mat, mat)

	//Threshold
	if threshtype == "mean"{
		gocv.AdaptiveThreshold(*mat, mat, (*th).MaxValue, 0 , 0, (*th).BlockSize, (*th).C)
	} else if threshtype == "gauss" {
		gocv.AdaptiveThreshold(*mat, mat, (*th).MaxValue, 1 , 0, (*th).BlockSize, (*th).C)
	} else {
		fmt.Println("Unavailable type")
	}

	//gocv.Threshold(*mat, mat, 0, (*th).MaxValue, 8)
}
