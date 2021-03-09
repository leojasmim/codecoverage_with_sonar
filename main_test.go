package main

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestEvenOrOdd_WhenValueEven_ThenResultShouldBeEven(t *testing.T) {
	value := 2

	result := evenOrOdd(value)
	
	assert.Equal(t, EVEN, result, "O resultado deve ser par")
}


func TestEvenOrOdd_WhenValueOdd_ThenResultShouldBeOdd(t *testing.T) {
	value := 1

	result := evenOrOdd(value)
	
	assert.Equal(t, ODD, result, "O resultado deve ser impar")
}