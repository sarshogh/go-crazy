package algorithms

import (
	"math"
	"fmt"
)

func IsPrime(number int) bool {
	number = int(math.Abs(float64(number)))

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimeNums(n int) []int {
	var result []int

	for i := 2; i < n; i++ {
		if IsPrime((i)) {
			// fmt.Println("prime number is:",i)
			result = append(result, i)
		}
	}

	return result
}
