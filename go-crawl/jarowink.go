package main

import (
	"fmt"
	"math"
)

func checkMatching(str1, str2 string) (float64, float64) {
	num := 0
	transpose := 0
	len1 := len(str1)
	matchArr1 := make([]bool, len1)
	len2 := len(str2)
	matchArr2 := make([]bool, len2)
	for i := range str1 {
		for j := range str2 {
			if str1[i] == str2[j] {
				distance := math.Abs(float64(i - j))
				allowedDistance := math.Floor(math.Max(float64(len1), float64(len2))/2) - 1
				if distance > allowedDistance {
					continue
				} else {
					matchArr1[i] = true
					matchArr2[j] = true
					num++
				}
			}
		}
	}
	// FIXME: Broken transposition loop
	j := 0
	for i := range str1 {
		if matchArr1[i] {
			if matchArr2[j] {
				if str1[i] != str2[j] {
					transpose++
				}
			}
			j++
		}
	}
	return float64(num), float64(transpose / 2)
}

func jaroSimilarity(input, actual string) float64 {
	m, t := checkMatching(input, actual)
	len1 := float64(len(input))
	len2 := float64(len(actual))
	if m == 0 {
		return 0
	} else {
		sim := float64(((m / len1) + (m / len2) + ((m - t) / m)) / 3)
		return sim
	}
}

func main() {
	fmt.Println(jaroSimilarity("", ""))
}
