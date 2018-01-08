// transformation invariant image phash
// based on https://github.com/pippy360/transformationInvariantImageSearch
package main

import (
	"fmt"
	"log"
	"os"

	"net/http"
	_ "net/http/pprof"

	"github.com/azr/phash"
	"github.com/azr/phash/cmd"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println("Usage: dtc path/to/image.jpg")
		os.Exit(1)
	}
	go http.ListenAndServe(":6060", nil)
	img, _ := cmd.OpenImageFromPath(os.Args[1])
	{
		// if img.Bounds().Max.Y > 400 {
		// 	img = imaging.Resize(img, 0, 400, imaging.Lanczos)
		// }
		// if img.Bounds().Max.X > 400 {
		// 	img = imaging.Resize(img, 400, 0, imaging.Lanczos)
		// }
	}
	keypoints := phash.FindKeypoints(img, 1)

	triangles := keypoints.EveryTriangles(phash.EveryTrianglesOpts{
		LowerThreshold: 50,
		UpperThreshold: 500,
		MinArea:        50,
	})

	hashes := phash.GetImageHashesForTriangles(img, triangles)
	log.Println("hashes:", hashes)
}
