package cmd

import (
	"image"
	// preload all image decoders
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	// preload all image/x decoders
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

// OpenImageFromPath load an image from full path
// using image.Decode
func OpenImageFromPath(path string) (img image.Image, format string) {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatalf("%s. Current dir: %s", err, os.Args[0])
	}
	defer reader.Close()

	m, s, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return m, s
}
