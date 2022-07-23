package phash

import (
	"strconv"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		a    uint64
		b    uint64
		want int
	}{
		{parseBinary(knownLHiresHash), parseBinary(knownLHiresHash), 0},
		{0, parseBinary(knownLHiresHash), 31},
		{0, parseBinary("0000000000000000000000000000000000000000000000000000000000001000"), 1},
	}
	for n, tt := range tests {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			if got := Distance(tt.a, tt.b); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
