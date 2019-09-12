package util

import (
	"math"
)

const (
	ZIL = iota
	LI
	QA
)

func FromQa(qa float64, unit int, is_pack bool) float64 {
	rate := 1.0

	switch unit {
	case ZIL:
		rate = 1000000000000.0
	case LI:
		rate = 1000000.0
	}

	ret := qa / rate

	if is_pack {
		ret = math.Round(ret)
	}

	return ret
}

func ToQa(qa float64, unit int) float64 {
	rate := 1.0

	switch unit {
	case ZIL:
		rate = 1000000000000.0
	case LI:
		rate = 1000000.0
	}

	return qa * rate
}
