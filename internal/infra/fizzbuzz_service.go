package infra

import (
	"fizzbuzz_leboncoin_test/internal/domain/models"
	"fizzbuzz_leboncoin_test/internal/domain/services"
	"fmt"
	"strconv"
)

type fizzBuzzService struct{}

func NewFizzBuzzService() services.IFizzBuzzService {
	return &fizzBuzzService{}
}

func (fbSvc fizzBuzzService) FizzBuzz(params *models.FizzBuzz) []string {
	result := make([]string, 0)
	if params != nil && params.Fizz.MutlipleOf != 0 && params.Buzz.MutlipleOf != 0 {
		for i := 1; i <= params.Limit; i++ {
			switch {
			case i%(params.Fizz.MutlipleOf*params.Buzz.MutlipleOf) == 0:
				result = append(result, fmt.Sprintf("%v%v", params.Fizz.ReplaceWith, params.Buzz.ReplaceWith))
			case i%params.Fizz.MutlipleOf == 0:
				result = append(result, params.Fizz.ReplaceWith)
			case i%params.Buzz.MutlipleOf == 0:
				result = append(result, params.Buzz.ReplaceWith)
			default:
				result = append(result, strconv.Itoa(i))
			}
		}
	}

	return result
}
