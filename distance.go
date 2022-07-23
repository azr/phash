package phash

import "math/bits"

// Distance returns the distance between two hashes by calculating the
// number of different bits in the hash
func Distance(a, b uint64) int {
	hamming := a ^ b
	return bits.OnesCount64(hamming)
}
