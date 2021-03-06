package main

import "fmt"

func main() {
	var ints = map[string]int64{
		"first":  34,
		"second": 12,
	}

	var floats = map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf(
		"Non-Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		)
	)
}

func SumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
