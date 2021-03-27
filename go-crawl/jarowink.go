package main

import (
	"fmt"
	"math"
	"strings"
)

func checkMatching(str1 string, str2 string) int {
	num := 0
	strArr1 := strings.Split(str1, "")
	len1 := len(str1)
	strArr2 := strings.Split(str2, "")
	len2 := len(str2)
	for i, s := range strArr1 {
		for j, t := range strArr2 {
			if s == t {
				distance := math.Abs(float64(i - j))
				allowedDistance := math.Floor(math.Max(float64(len1), float64(len2))/2) - 1
				if distance > allowedDistance {
					continue
				} else {
					num++
				}
			}
		}
	}
	return num
}

func countTranspositions() {
	//FIXME: do transpositions things here, need to rewrite other functions
}

func jaroSimilarity(input string, actual string) float64 {
	m := float64(checkMatching(input, actual))
	len1 := float64(len(input))
	len2 := float64(len(actual))
	// TODO: Fix transposition calculation, need to count amount of matches which do not have the same index
	t := float64(m / 2)
	if m == 0 {
		fmt.Println("beep")
		return 0
	} else {
		sim := float64(((m / len1) + (m / len2) + ((m - t) / m)) / 3)
		return sim
	}
}

func main() {
	fmt.Println(jaroSimilarity("CRATE", "TRACE"))
}
