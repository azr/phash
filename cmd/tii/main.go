// transformation invariant image phash
// based on https://github.com/pippy360/transformationInvariantImageSearch
package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/azr/phash/geometry"

	"net/http"
	_ "net/http/pprof"

	"github.com/StephaneBunel/bresenham"
	"github.com/azr/gift"
	"github.com/azr/phash"
	"github.com/azr/phash/cmd"
	"github.com/azr/phash/geometry/triangle"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println("Usage: dtc path/to/image.jpg")
		os.Exit(1)
	}
	go http.ListenAndServe(":6060", nil)
	img, _ := cmd.OpenImageFromPath(os.Args[1])
	keypoints := phash.FindKeypoints(img)

	triangles := triangle.AllPossibilities(triangle.PossibilititesOpts{
		LowerThreshold: 50,
		UpperThreshold: 500,
		MinArea:        50,
	}, keypoints)

	size := geometry.Min(500, len(triangles))
	log.Printf("image has %d keypoints which resulted in %d valid triangles, keeping first %d", len(keypoints), len(triangles), size)
	triangles = triangles[:size]
	Debug(os.Args[1], keypoints, triangles, img)

	hashes := phash.FromTriangles(img, triangles)
	for _, hash := range hashes {
		print(hash, " ")
	}
}

// Debug outputs your image to a triangle.jpg image
// containing only triangle coordinates.
func Debug(path string, keypoints geometry.Points, ts []triangle.Triangle, img image.Image) {
	g := gift.New(
		gift.Mean(5, true),
	)

	// STUFF
	simplifiedImage := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(simplifiedImage, img)
	for _, p := range keypoints {
		simplifiedImage.Set(p.X, p.Y, color.RGBA{B: 255, R: 255})
		simplifiedImage.Set(p.X+1, p.Y+1, color.RGBA{B: 255, R: 255})
		simplifiedImage.Set(p.X-1, p.Y-1, color.RGBA{B: 255, R: 255})
		simplifiedImage.Set(p.X-1, p.Y+1, color.RGBA{B: 255, R: 255})
		simplifiedImage.Set(p.X+1, p.Y-1, color.RGBA{B: 255, R: 255})
	}

	for _, t := range ts {
		// draw red triangle
		bresenham.Bresenham(simplifiedImage, t[0].X, t[0].Y, t[1].X, t[1].Y, color.RGBA{R: 255})
		bresenham.Bresenham(simplifiedImage, t[2].X, t[2].Y, t[1].X, t[1].Y, color.RGBA{R: 255})
		bresenham.Bresenham(simplifiedImage, t[2].X, t[2].Y, t[0].X, t[0].Y, color.RGBA{R: 255})
	}

	cmd.WriteImageToPath(simplifiedImage, path+"full-triangle")
}
