package algorithms

// This is to check an input number is a prime number
func IsPrime(number int) bool {

	if number < 2 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// This is to generate prime numbers less than input number n. 
// 1< n <Int32
func GeneratePrimeNums(n int) []int {
	var result []int
	if n < 2 {
		panic("argument number must be bigger than 1")
	}
	for i := 2; i < n; i++ {
		if IsPrime((i)) {
			result = append(result, i)
		}
	}
	return result
}
