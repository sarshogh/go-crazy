package main

import (
	"testing"
)

func Test_IsPrime(t * testing.T)  {
	subject:= 5
	expected:= true
	result:= fns.IsPrime(subject)
	
	if(result == expected) {
		t.Log("the subject with value of %d is a prime number", subject)
	}else {
		t.Error("the number %d is not a prime number", subject)
	}
}