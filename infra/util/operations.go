package util

import (
	"errors"
	"math"
)

func Pearson(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("slices must be of equal length")
	}

	var abar, bbar float64
	var n int
	for i := range a {
		if !math.IsNaN(a[i]) && !math.IsNaN(b[i]) {
			abar += a[i]
			bbar += b[i]
			n++
		}
	}
	nf := float64(n)
	abar, bbar = abar/nf, bbar/nf

	var numerator float64
	var sumAA, sumBB float64

	for i := range a {
		if !math.IsNaN(a[i]) && !math.IsNaN(b[i]) {
			numerator += (a[i] - abar) * (b[i] - bbar)
			sumAA += (a[i] - abar) * (a[i] - abar)
			sumBB += (b[i] - bbar) * (b[i] - bbar)
		}
	}

	return numerator / (math.Sqrt(sumAA) * math.Sqrt(sumBB)), nil
}
