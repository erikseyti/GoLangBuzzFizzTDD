package main

import (
	"testing"
)

func TestFizzBuzzWithValu3MustReturnFizz(t *testing.T) {
	result := fizzBuzz(3)

	if result != "fizz"{
		t.Errorf("Expected Fizz but got: \n%v", result)
	}
}

func TestFizzBuzzWithValue5MustReturnBuzz(t *testing.T) {
	result := fizzBuzz(5)

	if result != "buzz"{
		t.Errorf("Expected Fizz but got: \n%v", result)
	}
}

func TestFizzBuzzWithValue2MustReturn2(t *testing.T) {
	result := fizzBuzz(2)

	if result != "2" {
		t.Errorf("Expected 2 but got: \n%v", result)
	}
}

func TestFizzBuzzWithMultipleOf3MustReturnFizz(t *testing.T) {
	result := fizzBuzz(6)
	if result != "fizz" {
		t.Errorf("Expected fizz but got: \n%v", result)
	}
}

func TestFizzBuzzWithMultipleOf5MustReturnFizz(t *testing.T) {
	result := fizzBuzz(10)
	if result != "buzz" {
		t.Errorf("Expected buzz but got: \n%v", result)
	}
}

func TestFizzBuzzWithMultipleOf5And3MustReturnFizzBuzz(t *testing.T) {
	result := fizzBuzz(15)
	if result != "fizzbuzz" {
		t.Errorf("Expected fizzbuzz but got: \n%v", result)
	}
}

func TestFizzBuzzWith300MustReturnFizzBuzz(t *testing.T) {
	result := fizzBuzz(300)
	if result != "fizzbuzz" {
		t.Errorf("Expected fizzbuzz but got: \n%v", result)
	}
}

