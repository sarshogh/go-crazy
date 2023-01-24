package algorithms

import (
	"fmt"
	"testing"
)

// table-driven tests
var sample_numbers = []struct {
	name   string
	input  int
	output bool
}{
	{"isPrime(1)", 1, true},
	{"isPrime(4)", 4, false},
	{"isPrime(21)", 21, false},
	{"isPrime(19)", 19, true},
}

func Test_Seve_IsPrime(t *testing.T) {
	for index, item := range sample_numbers {
		fmt.Printf("runnig IsPrime test case %d \n", index)
		t.Run(item.name, func(t *testing.T) {
			subject := item.input
			expected := item.output
			result := IsPrime(subject)
			if result != expected {
				t.Errorf("got %t, want %t", result, expected)

			}
		})
	}
}
