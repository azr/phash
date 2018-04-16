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
	log.Printf("keypoints: %d", len(keypoints))

	triangles := triangle.AllPossibilities(triangle.PossibilititesOpts{
		Src:                 img,
		LowerThresholdRatio: 0.00003,
		UpperThresholdRatio: 0.00008,
		MinAreaRatio:        0.00009,
	}, keypoints)
	log.Printf("triangles: %d", len(triangles))

	// n := 12
	// n = 2
	// if n >= len(triangles) {
	// 	n = len(triangles) / 2
	// }
	triangles[0].Debug(img)

	hashes := phash.GetImageHashesForTriangles(img, triangles)
	log.Println("hashes:", hashes)
}
