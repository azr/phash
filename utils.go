
package main

import (
    "image"
    "code.google.com/p/biogo.matrix"
)

func RoundingFactor(x float64) float64 {
    if ((x) >= 0) {
        return 0.5
    }
    return -0.5
}

func getImageMatrix(src image.Gray) (*matrix.Dense, error) {
    bounds := src.Bounds()
    mtx := make([][]float64, bounds.Max.X)
    for x := 0; x < bounds.Max.X; x++ {
        mtx[x] = make([]float64, bounds.Max.Y)
        for y := 0; y < bounds.Max.Y; y++ {
            _, _, b, _ := src.At(x, y).RGBA()
            mtx[x][y] = float64(b)
        }
    }
    return matrix.NewDense(mtx)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
