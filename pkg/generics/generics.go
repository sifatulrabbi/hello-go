package generics

func SumInts(m map[string]int64) (r int64) {
	for _, v := range m {
		r += v
	}
	return
}

func SumFloats(m map[string]float64) (r float64) {
	for _, v := range m {
		r += v
	}
	return
}

func SumIntOrFloats[K comparable, V int64 | float64](m map[K]V) (r V) {
	for _, v := range m {
		r += v
	}
	return
}

// Prints the first value of a map
func PrintFirstMapEl[V string | int64 | float64 | bool](m map[string]V) (r V) {
	indexes := []string{}
	for i := range m {
		indexes = append(indexes, i)
	}
	r = m[indexes[0]]
	return
}
