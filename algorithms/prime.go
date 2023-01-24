package algorithms

import (
	"math"
)

func IsPrime(number int) bool {
	number = int(math.Abs(float64(number)))

	for i:=2; i < number; i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}