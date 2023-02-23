package infra

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHit(t *testing.T) {
	statsService := NewStatsService()
	request := "some request"

	testCases := []struct {
		scenario     string
		params       []byte
		calltimes    int
		expectedHits int64
	}{
		{
			scenario:     "should not increment hits if request is empty",
			params:       []byte(nil),
			calltimes:    5,
			expectedHits: 0,
		},
		{
			scenario:     "should increment hits each time called",
			params:       []byte(request),
			calltimes:    5,
			expectedHits: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.scenario, func(t *testing.T) {
			for i := 0; i < testCase.calltimes; i++ {
				statsService.Hit(testCase.params)
			}
			actualRequest, actualTotalHits := statsService.GetTopRequest()
			assert.Equal(t, testCase.expectedHits, actualTotalHits)
			if actualTotalHits > int64(0) {
				assert.Equal(t, request, string(actualRequest))
			}
		})
	}
}
