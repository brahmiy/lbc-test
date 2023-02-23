package models

type FizzBuzz struct {
	Limit     int
	Fizz      ReplaceMultipleOfWith
	Buzz      ReplaceMultipleOfWith
}

type ReplaceMultipleOfWith struct {
	MutlipleOf  int
	ReplaceWith string
}
