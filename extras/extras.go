package extras

import (
	"fmt"
	"image"


	"gocv.io/x/gocv"
)


func MaximizeContrast(mat *gocv.Mat, ans *gocv.Mat) {
	var imgTopHat gocv.Mat
	imgTopHat = gocv.NewMat()
	defer imgTopHat.Close()
	var imgBlackHat gocv.Mat
	imgBlackHat = gocv.NewMat()
	defer imgBlackHat.Close()
	var imgGrayscalePlusTopHat gocv.Mat
	imgGrayscalePlusTopHat = gocv.NewMat()
	defer imgGrayscalePlusTopHat.Close()
	var structuringElement gocv.Mat
	structuringElement = gocv.GetStructuringElement(gocv.MorphRect, image.Point{3, 3})
	gocv.MorphologyEx(*mat, &imgTopHat, gocv.MorphTophat, structuringElement)
	gocv.MorphologyEx(*mat, &imgBlackHat, gocv.MorphBlackhat, structuringElement)
	gocv.Add(*mat, imgTopHat, &imgGrayscalePlusTopHat)
	gocv.Subtract(imgGrayscalePlusTopHat, imgBlackHat, ans)
}



func PercentBlackWhite(src *gocv.Mat) []float64 {
	data := src.ToBytes()
	b := 0;
	w := 0;
	for i := 0; i < src.Rows(); i++ {
		for j := 0; j < src.Cols(); j++ {
			//fmt.Printf("src[%d][%d] = %d\n", i, j, data[j + i * src.Cols()])
			if data[j + i * src.Cols()] == 255 {
				w++
			}
			if data[j + i * src.Cols()] == 0 {
				b++
			}
		}
	}
	white := 100 * float64(w) / float64(src.Rows()) / float64(src.Cols())
	black := 100 * float64(b) / float64(src.Rows()) / float64(src.Cols())
	return []float64{white, black}
}

func DeleteBorders(src gocv.Mat) gocv.Mat {
	data := src.ToBytes()
	for i := 0; i < src.Rows(); i++ {
		for j := 0; j <= 10; j++{
			data[j + i * src.Cols()] = 255 
		}
		for j := src.Cols() - 10; j <= src.Cols() -1; j++ {
			data[j + i * src.Cols()] = 255
		}
	}
	for j := 0; j < src.Cols(); j++ {
		for i := 0; i <= 10; i++ {
			data[j + i * src.Cols()] = 255 
		}
		for i := src.Rows() - 10; i <= src.Rows() - 1; i++ {
			data[j + i * src.Cols()] = 255
		}
	}
	src, err := gocv.NewMatFromBytes(src.Rows(), src.Cols(), gocv.MatTypeCV8U, data)
	fmt.Println(err)
	return src
}
