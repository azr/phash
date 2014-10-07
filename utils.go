package phash

import ()

// HammingDistance computes hamming distance between two dtc hashes
func HammingDistance(hash1, hash2 uint64) uint64 {
	x := hash1 ^ hash2
	var m1, m2, h01, m4 uint64 = 0x5555555555555555, 0x3333333333333333, 0x0101010101010101, 0x0f0f0f0f0f0f0f0f
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}
