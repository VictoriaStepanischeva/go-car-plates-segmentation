package layers

import (
	"image/color"
	"image"
	"gocv.io/x/gocv"
)

type DrawingContours struct {
	Scale float64
	ScaleHWMin float64
	ScaleHWMax float64
	HMax float64
	WMax float64
}

func NewDrawContours(scale float64, scalehwmin float64, scalehwmax float64, hmax float64, wmax float64) DrawingContours{
	return DrawingContours{
		Scale: scale,
		ScaleHWMin: scalehwmin,
		ScaleHWMax: scalehwmax,
		HMax: hmax,
		WMax: wmax,
	}
}

func (dc *DrawingContours) FindAndDraw (mat *gocv.Mat, matrgb *gocv.Mat) {
	res := gocv.FindContours(*mat, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	for j := 0; j < len(res); j++ {
		rectangle := gocv.BoundingRect(res[j])
		scaley := float64(rectangle.Dy())/float64(mat.Rows())
		scalehw := float64(rectangle.Dy())/float64(rectangle.Dx())
		if (scalehw > (*dc).ScaleHWMin) && (scalehw <= (*dc).ScaleHWMax) && (scaley >= (*dc).Scale) && (scaley <= (*dc).Scale + 0.55) && (float64(rectangle.Dy()) >= (*dc).HMax) && (float64(rectangle.Dx()) >= (*dc).WMax) {
			//another = append(another, res[j])
			gocv.Rectangle(matrgb, rectangle,  color.RGBA{0, 255, 0, 1}, 2)
		}
	}
}

func (dc *DrawingContours) FindAndDrawContours (mat *gocv.Mat, matrgb *gocv.Mat) {
	another := make([][]image.Point, 0)
	res := gocv.FindContours(*mat, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	for j := 0; j < len(res); j++ {
		rectangle := gocv.BoundingRect(res[j])
		scaley := float64(rectangle.Dy())/float64(mat.Rows())
		scalehw := float64(rectangle.Dy())/float64(rectangle.Dx())
		if (scalehw > (*dc).ScaleHWMin) && (scalehw <= (*dc).ScaleHWMax) && (scaley >= (*dc).Scale) && (scaley <= (*dc).Scale + 0.55) && (float64(rectangle.Dy()) >= (*dc).HMax) && (float64(rectangle.Dx()) >= (*dc).WMax) {
			another = append(another, res[j])
		}
	}
	gocv.DrawContours(matrgb, another, -1, color.RGBA{0, 255, 0, 1}, 1)
}

func (dc *DrawingContours) FindAndDrawAllContours (mat *gocv.Mat, matrgb *gocv.Mat) {
	res := gocv.FindContours(*mat, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	gocv.DrawContours(matrgb, res, -1, color.RGBA{0, 255, 0, 1}, 1)
}

func (dc *DrawingContours) FindContoursCanny (mat *gocv.Mat) {
	gocv.Canny(*mat, mat, 100, 255)
}

