package main

import (
    "io/ioutil"
    "testing"
    // "time"
)

// var images_dir = "./tests/cats/"

func loadImages() []ImageBag {
    images_in_dir, err := ioutil.ReadDir(images_dir)
    if err != nil {
        panic(err)
    }

    images := make([]ImageBag, len(images_in_dir) )
    for i, fi := range images_in_dir {
        images[i].InitialiseFromFileInfo(fi)
    }
    return images
}

// func TestTimeConsuming(t *testing.T) {
//     if testing.Short() {
//         t.Skip("skipping test in short mode.")
//     }

// }

func BenchmarkDctImageHash_one(b *testing.B) {

    images := loadImages()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        for _, img := range images {
            img.ComputeImageHashOne()
        }
    }

}

func BenchmarkDctImageHash_two(b *testing.B) {

    images := loadImages()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        for _, img := range images {
            img.ComputeImageHashTwo()
        }
    }

}

func BenchmarkDctImageHash_phash(b *testing.B) {

    images := loadImages()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        for _, img := range images {
            img.ComputeImageHashPhash()
        }
    }

}
