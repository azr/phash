package phash

import (
	"go/build"
	"os"
	"strconv"
)

func parseBinary(i string) uint64 {
	v, _ := strconv.ParseInt(i, 2, 64)
	return uint64(v)
}

type Empty struct{}

func gopath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}
