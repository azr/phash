package phash

// copy pasta from https://ironchef-team21.googlecode.com/git-history/75856e07bb89645d0e56820d6e79f8219a06bfb7/ironchef_team21/src/ImagePHash.java

import (
	"image"
	"math"
	"sort"

	"github.com/disintegration/imaging"
)

var (
	DTCSizeBig = 32
	DTCSize    = 8
)

// DTC computes perceptual hash for image
// using phash dtc image technique
func DTC(img image.Image) (phash uint64) {
	if img == nil {
		return
	}

	size := DTCSizeBig
	smallerSize := DTCSize

	/* 1. Reduce size.
	 * Like Average Hash, pHash starts with a small image.
	 * However, the image is larger than 8x8; 32x32 is a good size.
	 * This is really done to simplify the DCT computation and not
	 * because it is needed to reduce the high frequencies.
	 */
	im := imaging.Resize(img, size, size, imaging.Lanczos)

	/* 2. Reduce color.
	 * The image is reduced to a grayscale just to further simplify
	 * the number of computations.
	 */

	vals := make([]float64, size*size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			r, g, b, _ := im.At(i, j).RGBA()
			vals[size*i+j] = 0.299*float64(r) +
				0.587*float64(g) +
				0.114*float64(b)
		}
	}

	/* 3. Compute the DCT.
	 * The DCT separates the image into a collection of frequencies
	 * and scalars. While JPEG uses an 8x8 DCT, this algorithm uses
	 * a 32x32 DCT.
	 */

	applyDCT2 := func(N int, f []float64) []float64 {
		// initialize coefficients
		c := make([]float64, N)
		c[0] = 1 / math.Sqrt(2)
		for i := 1; i < N; i++ {
			c[i] = 1
		}

		// output goes here
		F := make([]float64, N*N)

		// construct a lookup table, because it's O(n^4)
		entries := (2 * N) * (N - 1)
		COS := make([]float64, entries)
		for i := 0; i < entries; i++ {
			COS[i] = math.Cos(float64(i) / float64(2*N) * math.Pi)
		}

		// the core loop inside a loop inside a loop...
		for u := 0; u < N; u++ {
			for v := 0; v < N; v++ {
				var sum float64
				for i := 0; i < N; i++ {
					for j := 0; j < N; j++ {
						sum += COS[(2*i+1)*u] *
							COS[(2*j+1)*v] *
							f[N*i+j]
					}
				}
				sum *= ((c[u] * c[v]) / 4)
				F[N*u+v] = sum
			}
		}
		return F
	}

	dctVals := applyDCT2(size, vals)

	/* 4. Reduce the DCT.
	 * This is the magic step. While the DCT is 32x32, just keep the
	 * top-left 8x8. Those represent the lowest frequencies in the
	 * picture.
	 */

	vals = make([]float64, 0, smallerSize*smallerSize)
	for x := 1; x <= smallerSize; x++ {
		for y := 1; y <= smallerSize; y++ {
			vals = append(vals, dctVals[size*x+y])
		}
	}

	/* 5. Compute the average value.
	 * Like the Average Hash, compute the mean DCT value (using only
	 * the 8x8 DCT low-frequency values and excluding the first term
	 * since the DC coefficient can be significantly different from
	 * the other values and will throw off the average).
	 */

	sortedVals := make([]float64, smallerSize*smallerSize)
	copy(sortedVals, vals)
	sort.Float64s(sortedVals)
	median := sortedVals[smallerSize*smallerSize/2]

	/* 6. Further reduce the DCT.
	 * This is the magic step. Set the 64 hash bits to 0 or 1
	 * depending on whether each of the 64 DCT values is above or
	 * below the average value. The result doesn't tell us the
	 * actual low frequencies; it just tells us the very-rough
	 * relative scale of the frequencies to the mean. The result
	 * will not vary as long as the overall structure of the image
	 * remains the same; this will survive gamma and color histogram
	 * adjustments without a problem.
	 */

	for n, e := range vals {
		if e > median { // when frequency is higher than median
			phash ^= (1 << uint64(n)) // set nth bit to one
		}
	}
	return phash
}
