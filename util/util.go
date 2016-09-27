package util

import (
	"math"
)

func Minint64(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Round(a float64) int {
	return int(a + math.Copysign(0.5, a))
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func OctantFromDeltas(dX, dY int) int {
	if dX > 0 {
		if dY > 0 {
			if Abs(dX) > Abs(dY) {
				return 0
			} else {
				return 1
			}
		} else {
			if Abs(dX) > Abs(dY) {
				return 7
			} else {
				return 6
			}
		}
	} else {
		if dY > 0 {
			if Abs(dX) > Abs(dY) {
				return 3
			} else {
				return 2
			}
		} else {
			if Abs(dX) > Abs(dY) {
				return 4
			} else {
				return 5
			}
		}
	}
}

func SwitchToOctantZeroFrom(octant, x, y int) (int, int) {
	switch octant {
	case 0:
		return x, y
	case 1:
		return y, x
	case 2:
		return y, -x
	case 3:
		return -x, y
	case 4:
		return -x, -y
	case 5:
		return -y, -x
	case 6:
		return -y, x
	case 7:
		return x, -y
	default:
		return x, y
	}
}

func SwitchFromOctantZeroTo(octant, x, y int) (int, int) {
	switch octant {
	case 0:
		return x, y
	case 1:
		return y, x
	case 2:
		return -y, x
	case 3:
		return -x, y
	case 4:
		return -x, -y
	case 5:
		return -y, -x
	case 6:
		return y, -x
	case 7:
		return x, -y
	default:
		return x, y
	}
}
