package phash

import (
	"strconv"
)

func parseBinary(i string) uint64 {
	v, _ := strconv.ParseInt(i, 2, 64)
	return uint64(v)
}

type Empty struct{}
