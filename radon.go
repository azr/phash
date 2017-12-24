package phash

import (
	"image"
	"image/color"
	"math"
)

func addPixelsToGray(src image.Image, sx, sy int, dst image.Gray, dx, dy int) {
	clr := src.At(sx, sy)
	greyColor, _ := color.GrayModel.Convert(clr).(color.Gray)
	dst.Set(dx, dy, color.Gray{dst.At(dx, dy).(color.Gray).Y + greyColor.Y})
}

// RadonProjections computes radon projections of lines running through the image
// center for lines angled 0 to 180 degrees from horizontal.
//
//  /param img - img.Image src image
//  /param  N  - number :of angled lines to consider.
//  /return sinogram - Projections
//  /return nbPerLine - number of pixel per line in sinogram
func RadonProjections(src image.Image, N int) *image.Gray {

	// N, M := src.Bounds().Size().X, src.Bounds().Size().Y
	// n, m := N/2, M/2

	// // The total number of rho’s is the number of pixels on the diagonal, since
	// // this is the largest straight line on the image when rotating
	// rhomax := int(math.Ceil(math.Sqrt(float64(N*N + M*M))))
	// rc := rhomax / 2
	// mt := 180
	// imRadon := image.NewGray(image.Rect(0, 0, rhomax, mt))

	// for t := 0; t < 45; t++ { // bellow 45°, use y as a var
	// 	costheta := math.Cos(float64(t) * math.Pi / 180)
	// 	sintheta := math.Sin(float64(t) * math.Pi / 180)
	// 	a := -costheta / sintheta

	// 	for r := 0; r < rhomax; r++ {
	// 		rho := r - rc
	// 		b := float64(rho) / sintheta
	// 		ymax := math.Min(-a*float64(m)+b, float64(n-1))
	// 		ymin := math.Max(a*float64(m)+b, float64(-n))

	// 	}

	// }

	// /**
	//  * PURPOSE: perform radon transform of given image
	//  * PARAM: CImg<unsigned char> im - image to detect lines
	//  * 		  N :- number of angles to consider (should be a power of 2)
	//  *                (the values of N will be spread over 0 to 2PI)
	//  * RETURN CImg<unsigned char> - transform of given image of size, N x D
	//  *                              D = rhomax = sqrt(dimx*dimx + dimy*dimy)/2
	//  **/

	size := src.Bounds().Size()
	image_height, image_width := size.X, size.Y

	//calc offsets to center the image
	xofftemp := float64(image_width)/2.0 - 1
	yofftemp := float64(image_height)/2.0 - 1
	xoffset := int(math.Floor(xofftemp + roundingFactor(xofftemp)))
	yoffset := int(math.Floor(yofftemp + roundingFactor(yofftemp)))
	dtemp := math.Sqrt(float64(xoffset*xoffset + yoffset*yoffset))
	D := int(math.Floor(dtemp + roundingFactor(dtemp)))

	imRadon := image.NewGray(image.Rect(0, 0, N, D))

	//for each angle k to consider
	for k := 0; k < N; k++ {
		//only consider from PI/8 to 3PI/8 and 5PI/8 to 7PI/8
		//to avoid computational complexity of a steep angle
		if k == 0 {
			k = N / 8
			continue
		} else if k == (3*N/8 + 1) {
			k = 5 * N / 8
			continue
		} else if k == 7*N/8+1 {
			k = N
			continue
		}

		//for each rho length, determine linear equation and sum the line
		//sum is to sum the values along the line at angle k2pi/N
		//sum2 is to sum the values along the line at angle k2pi/N + N/4
		//The sum2 is performed merely by swapping the x,y axis as if the image were rotated 90 degrees.
		for d := 0; d < D; d++ {
			theta := 2 * float64(k) * math.Pi / float64(N)                              //calculate actual theta
			alpha := math.Tan(theta + math.Pi/2)                                        //calculate the slope
			beta_temp := -alpha*float64(d)*math.Cos(theta) + float64(d)*math.Sin(theta) //y-axis intercept for the line
			beta := int(math.Floor(beta_temp + roundingFactor(beta_temp)))
			//for each value of m along x-axis, calculate y
			//if the x,y location is within the boundary for the respective image orientations, add to the sum
			var sum1, sum2 float64
			M := max(image_width, image_height)
			for m := 0; m < M; m++ {
				//interpolate in-between values using nearest-neighbor approximation
				//using m,n as x,y indices into image
				n_temp := alpha*float64(m-xoffset) + float64(beta)
				n := int(math.Floor(n_temp + roundingFactor(n_temp)))
				if (m < image_width) && (n+yoffset >= 0) && (n+yoffset < image_height) {
					c := src.At(m, n+yoffset)
					// v :=
					sum1 += float64(color.GrayModel.Convert(c).(color.Gray).Y) / 130
				}
				n_temp = alpha*float64(m-yoffset) + float64(beta)
				n = int(math.Floor(n_temp + roundingFactor(n_temp)))
				if (m < image_height) && (n+xoffset >= 0) && (n+xoffset < image_width) {
					c := src.At(-(n+xoffset)+image_width-1, m)
					// v :=
					sum2 += float64(color.GrayModel.Convert(c).(color.Gray).Y) / 130
				}
			}
			//assign the sums into the result matrix
			imRadon.Set(k, d, color.Gray{uint8(sum1)})
			//assign sum2 to angle position for theta + PI/4
			imRadon.Set(((k + N/4) % N), d, color.Gray{uint8(sum2)})
		}
	}
	return imRadon

	// Steps = 360
	// size := src.Bounds().Size()
	// height, width := size.X, size.Y

	// D := math.Sqrt(height*height+width*width) / 2
	// xCenter, yCenter := float64(width)/2.0, float64(height)/2.0
	// xOff, yOff := int(math.Floor(xCenter)), int(math.Floor(yCenter))

	// radonMap := image.NewGray(image.Rect(0, 0, Steps, D))

	// addColorsToGreyScale := func(a, b color.Color) color.Color {
	// 	a2 := color.GrayModel.Convert(a).(color.Gray)
	// 	a2.Y = color.GrayModel.Convert(b).(color.Gray).Y
	// 	return a2
	// }
	// for k := 0 k < Steps k++ {

	// 	for x := 0 x < D x++ {
	// 		X := x - xOff

	// 	}
	// }
	// return radonMap

	// size := src.Bounds().Size()
	// out := image.NewGray(image.Rect(0, 0, Steps, max(size.X, size.Y)))
	// for s := 0 s < Steps s++ {
	// 	θ := -float64(s) * 180 / float64(Steps)
	// 	rotation := imaging.Rotate(src, θ, color.Black)
	// 	for i, v := range sum(rotation) {
	// 		out.Set(s, i, color.Gray{uint8(v / float64(size.Y))})
	// 	}
	// }
	// return out
	// step := 180.0 / float64(N)

	// size := src.Bounds().Size()
	// overX := int(float64(size.X) * 2)
	// overY := int(float64(size.Y) * 2)
	// var img image.Image = image.NewGray(image.Rect(0, 0, size.X+overX, size.Y+overY))
	// img = imaging.Overlay(img, src, image.Pt(overX/2, overY/2), 1)
	// size = img.Bounds().Size()

	// D := max(size.X, size.Y)
	// out := image.NewGray(image.Rect(0, 0, N, D))

	// // imaging.Rotate uses paralellisation
	// // based on runtime.GOMAXPROCS(0),
	// // so we set it to 1 during that period
	// // this paralellisation is not usefull
	// numProcs := runtime.GOMAXPROCS(0)
	// runtime.GOMAXPROCS(1)
	// defer func() {
	// 	runtime.GOMAXPROCS(numProcs)
	// }()

	// // parallel(N, func(n, N int) {
	// // for each given angle θ
	// for n := 0 n < N n++ {
	// 	θ := float64(n) * step

	// 	//have a duplicate img rotated by θ
	// 	draw := imaging.Rotate(img, θ, black)

	// 	sinogram := make([]float64, D)
	// 	// get column average profile
	// 	for y := 0 y < size.Y y++ {
	// 		for x := 0 x < size.X x++ {
	// 			greyColor, _ := color.GrayModel.Convert(draw.At(x, y)).(color.Gray)
	// 			sinogram[x] = sinogram[x] + float64(greyColor.Y)
	// 		}
	// 	}

	// 	//Set out line with sinogram
	// 	for d := 0 d < D d++ {
	// 		out.Set(n, d, color.Gray{uint8(sinogram[d] / float64(size.Y))})
	// 	}
	// }
	// // })

	// return out

	// parallel(N, func(n, N int) {
	// for each given angle θ
	// D := max(size.Y, size.X)
	// out := image.NewGray(image.Rect(0, 0, Steps, D))
	// for n := 0 n < Steps n++ {
	// 	θ := -float64(n) * 180 / float64(Steps)

	// 	sinogram := make([]float64, D)
	// 	// get column average profile
	// 	for y := 0 y < size.Y y++ {
	// 		for x := 0 x < size.X x++ {
	// 			x*math.Cos(θ) + y*math.Sin(θ)
	// 			sinogram[x] = sinogram[x]
	// 		}
	// 	}

	// 	//Set out line with sinogram
	// 	for d := 0 d < D d++ {
	// 		out.Set(n, d, color.Gray{uint8(sinogram[d] / float64(size.Y))})
	// 	}
	// }
	// })

	// return out
}
