package infra

import (
	"fizzbuzz_leboncoin_test/internal/domain/services"
	"fmt"
	"sync"
)

type statsService struct {
	// use of concurrent-safe map for stats
	stats *sync.Map
}

func NewStatsService() services.IStatsService {
	return &statsService{
		stats: &sync.Map{},
	}
}

func (statsService *statsService) Hit(request []byte) {
	if len(request) > 0 {
		updatedHit := int64(1)
		hitCount := statsService.getTotalHits(request)
		if hitCount != nil {
			updatedHit += *hitCount
		}
		statsService.stats.Store(string(request), updatedHit)
	}
}

func (statsService *statsService) GetTopRequest() ([]byte, int64) {
	topRequest, totalHits := statsService.getTopRequest()
	if topRequest != nil && totalHits != nil {
		return []byte(*topRequest), *totalHits
	}
	return nil, 0
}


func (statsService *statsService) getTotalHits(request []byte) *int64 {
	count, exists := statsService.stats.Load(string(request))
	if exists {
		value := count.(int64)
		return &value
	}
	return nil
}

func (statsService *statsService) getTopRequest() (*string, *int64) {
	var topRequest *string
	var hitCount int64

	statsService.stats.Range(func(k, v interface{}) bool {
		if v.(int64) > hitCount {
			parsedRequest := fmt.Sprintf("%v", k)
			topRequest = &parsedRequest
			hitCount = v.(int64)
		}
		return true
	})

	return topRequest, &hitCount
}
