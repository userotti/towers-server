package utility

import "time"

// linear does a interpolation returns [0..1] floats
func linear(t0 time.Time, t1 time.Time, t time.Time) float32 {
	if t.After(t1) || t.Equal(t1) {
		return 1
	}

	if t.Before(t0) || t.Equal(t0) {
		return 0
	}

	return float32(float32(t.Sub(t0).Nanoseconds()) / float32(t1.Sub(t0).Nanoseconds()))
}
