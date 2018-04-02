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
	keypoints := phash.FindKeypoints(img)
	log.Printf("keypoints: %d", len(keypoints))

	bounds := img.Bounds()
	triangles := keypoints.EveryTriangles(phash.EveryTrianglesOpts{
		Pixels:              (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y),
		LowerThresholdRatio: 0.00003,
		UpperThresholdRatio: 0.00008,
		MinAreaRatio:        0.00009,
	})
	log.Printf("triangles: %d", len(triangles))

	n := 12
	// n = 2
	// if n >= len(triangles) {
	// 	n = len(triangles) / 2
	// }
	triangles[n].Debug(img)

	hashes := phash.GetImageHashesForTriangles(img, triangles)
	log.Println("hashes:", hashes)
}
