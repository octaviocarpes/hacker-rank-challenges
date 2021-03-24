package main

import (
	"fmt"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func findNextMultipleOfFive(grade int32) int32 {
	var counter int32 = 1
	var flag bool = true

	for flag {
		if (grade + counter) % 5 != 0 {
			counter++
		} else {
			flag = false
		}
	}

	return grade + counter
}

func gradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		grade := grades[i]
		if grade < 38 {
			continue
		} else {
			nextMultiple := findNextMultipleOfFive(grade)
			fmt.Println("Next multiple of", grade, "is", nextMultiple)
			if (nextMultiple - grade) < 3 {
				grades[i] = nextMultiple
			} else {
				continue
			}
		}
	}

	fmt.Println(grades)
	return grades
}

func main() {
	grades := []int32{73, 67, 38, 33}

	fmt.Println(gradingStudents(grades))
}

