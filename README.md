phash
![Coverage](https://img.shields.io/badge/Coverage-54.3%25-yellow)
![ci.yml](https://github.com/azr/phash/actions/workflows/ci.yml/badge.svg)
=====

A simple perceptual hashing library in Go.
Only works for images.

Usage :

```go
f, err := os.Open("image.jpg")
if err != nil {
    panic(err)
}
defer f.Close()
img, _, err := image.Decode(f)
if err != nil {
    panic(err)
}
phash.DTC(img)
```