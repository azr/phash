package cmd

import (
	"image"
	// preload all image decoders
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"

	// preload all image/x decoders
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

// OpenImageFromPath load an image from full path
// using image.Decode
// or panics
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

// WriteImageToPath creates an image somewhere for you
// or panics
func WriteImageToPath(img image.Image, path string) {
	f, err := os.Create(path + ".jpg")
	if err != nil {
		log.Fatalf("%s. Current dir: %s", err, os.Args[0])
	}
	defer f.Close()
	err = jpeg.Encode(f, img, nil)
	if err != nil {
		log.Fatalf("Write err: %s. Current dir: %s", err, os.Args[0])
	}
}
