package main

import (
	"fmt"
	"sort"
)

//Sorting slices // ordenando slices en orden alfabetico o numérico

func main() {

	sInts := []int{1, 0, 2, -3, 4, -20}
	fmt.Println(sInts)
	sort.Ints(sInts)
	fmt.Println(sInts)

	sFloats := []float64{1.0, 0.2, 0.22, -3, 4.1, -0.1}
	fmt.Println(sFloats)
	sort.Float64s(sFloats)
	fmt.Println(sFloats)

	sStrings := []string{"aa", "a", "A", "Aa", "aab", "AAa"}
	fmt.Println(sStrings)
	sort.Strings(sStrings)
	fmt.Println(sStrings)

}
