package main

import (
	"fmt"
)

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var fallenApples int32 = 0
	var fallenOranges int32 = 0
	var houseStartingPoint int32 = s
	var houseEndingPoint int32 = t
	var appleTreePoint int32 = a
	var orangeTreePoint int32 = b

	for i := 0; i < len(apples); i++ {
		applePoint := apples[i]
		if applePoint + appleTreePoint >= houseStartingPoint && applePoint + appleTreePoint <= houseEndingPoint {
			fallenApples++
		}
	}

	for i := 0; i < len(oranges); i++ {
		orangePoint := oranges[i]
		if orangePoint + orangeTreePoint >= houseStartingPoint && orangePoint + orangeTreePoint <= houseEndingPoint {
			fallenOranges++
		}
	}

	fmt.Println(fallenApples)
	fmt.Println(fallenOranges)
}

func main() {
	var houseStartingPoint int32 = 7
	var houseEndingPoint int32 = 11

	var appleTreePoint int32 = 5
	var orangeTreePoint int32 = 15

	fallenApplesPoints := []int32{-2, 2, 1}
	fallenOrangesPoints := []int32{5, -6}

	countApplesAndOranges(
		houseStartingPoint,
		houseEndingPoint,
		appleTreePoint,
		orangeTreePoint,
		fallenApplesPoints,
		fallenOrangesPoints,
	)
}
