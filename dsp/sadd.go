// +build amd64

package dsp

// Mixes together two audio frames in L16 format.
// Uses saturating addition.
func Sadd(dst, src []int16) {
	n := len(dst)
	if len(src) < n {
		n = len(src)
	}
	if n == 0 {
		return
	}

	sadd(&dst[0], &src[0], n)
}

func sadd(dst, src *int16, n int)
