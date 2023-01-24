package algorithms

import (
	"testing"
)

func Test_IsPrime(t *testing.T) {
	subject := 5
	expected := true
	result := IsPrime(subject)

	if result == expected {
		t.Logf("the subject with value of %d is a prime number", subject)
	} else {
		t.Errorf("the number %d is not a prime number", subject)
	}
}
