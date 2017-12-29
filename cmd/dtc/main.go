package main

import (
	"fmt"
	"os"

	"github.com/azr/phash"
	"github.com/azr/phash/cmd"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println("Usage: dtc path/to/image.jpg")
		os.Exit(1)
	}
	img, _ := cmd.OpenImageFromPath(os.Args[1])
	fmt.Printf("%d\n", phash.DTC(img))
}
