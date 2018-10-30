package main

import (
	"log"
	"fmt"
	"image"
	"flag"
	"path/filepath"
	"time"
	"detect/extras"
	"detect/layers"
	"io/ioutil"

	"gocv.io/x/gocv"
)

func main() {

	//Command-line args
	var srcd, dstd, pattern string
	flag.StringVar(&srcd, "srcd", ".", "directory with source images")
	flag.StringVar(&dstd, "dstd", "/tmp", "directory with destination images")
	flag.StringVar(&pattern, "pattern", "*", "pattern for source files")
	flag.Parse()
	if len(flag.Args()) != 0 {
		log.Fatal("Invalid args, please read help")
	}
	fmt.Println("srcd:", srcd)
	fmt.Println("dstd:", dstd)
	fmt.Println("pattern:", pattern)
	files, err := ioutil.ReadDir(srcd)
	if err != nil {
		log.Fatal(err)
	}
	ar := make([]string, 0)
	rr := make([]string, 0)
	for _, f := range files {
		matched, err := filepath.Match(pattern, f.Name())
		if err != nil {
			log.Fatal(err)
		}
		if matched {
			ar = append(ar, filepath.Join(srcd, f.Name()))
			rr = append(rr, filepath.Join(dstd, f.Name()))
		}
	}
	
	//Processing
	for i := 0; i < len(ar); i++{
		src := gocv.IMRead(ar[i], gocv.IMReadColor)
		src_b := gocv.NewMat()
		
		start := time.Now()
		//Resize
		gocv.Resize(src, &src, image.Point { 520, 112 }, 0, 0, 2);
		gocv.Resize(src, &src_b, image.Point { 520, 112 }, 0, 0, 2);
		
		//Maximize contras using morphology
		extras.MaximizeContrast(&src, &src)
		
		//Unsharping using Gauss
		srcunsh := src
		nsunsh := layers.NewNoiseSuppression()
		lstunsh := []string{"gauss"}
		kernelGaussunsh := image.Pt(3, 3)
		sigmaGaussunsh := 3.0
		nsunsh.AddGauss(kernelGaussunsh, sigmaGaussunsh)
		nsunsh.Supress(&srcunsh, lstunsh)
		gocv.AddWeighted(src, 1.25, srcunsh, -0.25, 0, &src)

		//NoiseSuppression using Gauss and Median
		ns := layers.NewNoiseSuppression()
		lst := []string{"gauss", "median"}
		kernelGauss := image.Pt(3, 3)
		sigmaGauss := 3.0
		kernelMedian := 3
		ns.AddGauss(kernelGauss, sigmaGauss)
		ns.AddMedian(kernelMedian)
		ns.Supress(&src, lst)

		//Threshold (RGBToGrayScale, Equlization, AdaptiveThreshold)
		th := layers.NewThreshold(255, 25, 14)
		th.Binarize(&src, "mean")
		

		//Morhology (dilate until white is equal determined value)
		res := extras.PercentBlackWhite(&src)
		for res[0] <= 69.0 {
			structuringElement := gocv.GetStructuringElement(gocv.MorphDilate, image.Point{3, 3})
			gocv.Dilate(src, &src,structuringElement)
			res = extras.PercentBlackWhite(&src)
		}

		//Delete black borders
		src = extras.DeleteBorders(src)
		//qq0 := fmt.Sprintf("%s/%d_tt.png", dstd, i)
		//gocv.IMWrite(qq0, src)
		
		//Finding and Drawing Contours
		scale := 0.3
		scaleHWMin := 0.5
		scaleHWMax := 3.0
		HMin := 27.0
		WMin := 27.0
		dc := layers.NewDrawContours(scale, scaleHWMin, scaleHWMax, HMin, WMin)
		dc.FindContoursCanny(&src)
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println(elapsed)
		dc.FindAndDraw(&src, &src_b)

		gocv.IMWrite(rr[i], src_b)
	}	
}
