package bootstrap

import (
	"fizzbuzz_leboncoin_test/internal/infra"
	"fizzbuzz_leboncoin_test/internal/ui"
	"fmt"
)

func LaunchFizzBuzz() {

	fizzBuzzService := infra.NewFizzBuzzService()
	statsService := infra.NewStatsService()
	fizzBuzzHandler := ui.NewFizzBuzzHandler(fizzBuzzService, statsService)

	err := ui.ServeRouter(fizzBuzzHandler)
	if err != nil {
		panic(fmt.Errorf("an error happened while serving the app: %w", err))
	}
}
