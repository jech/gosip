// +build !amd64
// +build !arm64

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

	for i, v := range src[:n] {
		w := int32(dst[i]) + int32(v)
		if w >= 0x7FFF {
			w = 0x7FFF
		}
		if w < -0x8000 {
			w = -0x8000
		}
		dst[i] = int16(w)
	}
}

