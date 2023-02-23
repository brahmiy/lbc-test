package services

import "fizzbuzz_leboncoin_test/internal/domain/models"

type IFizzBuzzService interface {
	FizzBuzz(params *models.FizzBuzz) []string
}
