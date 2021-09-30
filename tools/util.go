package editor

import "strconv"

func parseI(i string, d int) int {
	o, err := strconv.Atoi(i)
	if err != nil {
		o = d
	}

	return o
}

func parseF(f string, d float64) float64 {
	o, err := strconv.ParseFloat(f, 64)
	if err != nil {
		o = d
	}

	return o
}
