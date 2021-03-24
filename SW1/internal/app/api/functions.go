package api

import (
	"SW1/internal/app/models"
	"math"
)

func findRoots(c models.Coefficients) int {
	if c.A == 0 && c.B == 0 {
		return 0
	}

	if c.A == 0 {
		return 1
	}

	d := int(math.Pow(float64(c.B), 2)) - 4*c.A*c.C

	if d > 0 {
		return 2
	}

	if d == 0 {
		return 1
	}

	return 0
}
