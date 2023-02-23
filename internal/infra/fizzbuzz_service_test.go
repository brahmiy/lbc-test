package infra

import (
	"fizzbuzz_leboncoin_test/internal/domain/models"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFizzBuzz(t *testing.T) {
	fizzBuzzService := NewFizzBuzzService()
	testCases := []struct {
		scenario string
		params   *models.FizzBuzz
		expected []string
	}{
		{
			scenario: "should return empty array if params is nil",
			params:   nil,
			expected: []string{},
		},
		{
			scenario: "should return empty array if empty params",
			params:   &models.FizzBuzz{},
			expected: []string{},
		},
		{
			scenario: "should return correct array",
			params: &models.FizzBuzz{
				Limit: 11,
				Fizz: models.ReplaceMultipleOfWith{
					MutlipleOf:  3,
					ReplaceWith: "fizz",
				},
				Buzz: models.ReplaceMultipleOfWith{
					MutlipleOf:  5,
					ReplaceWith: "buzz",
				},
			},
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11"},
		},
		{
			scenario: "should return correct array when limit is above multiple of given values",
			params: &models.FizzBuzz{
				Limit: 16,
				Fizz: models.ReplaceMultipleOfWith{
					MutlipleOf:  3,
					ReplaceWith: "fizz",
				},
				Buzz: models.ReplaceMultipleOfWith{
					MutlipleOf:  5,
					ReplaceWith: "buzz",
				},
			},
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.scenario, func(t *testing.T) {
			actual := fizzBuzzService.FizzBuzz(testCase.params)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
