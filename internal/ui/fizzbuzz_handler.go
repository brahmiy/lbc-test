package ui

import (
	"encoding/json"
	"fizzbuzz_leboncoin_test/internal/domain/models"
	"fizzbuzz_leboncoin_test/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type FizzBuzzHandler interface {
	handleFizzBuzz(ctx *gin.Context)
	handleTopRequests(ctx *gin.Context)
}

type fizzBuzzHandler struct {
	fizzBuzzService services.IFizzBuzzService
	statsService    services.IStatsService
}

func NewFizzBuzzHandler(fizzBuzzService services.IFizzBuzzService, statsService services.IStatsService) fizzBuzzHandler {
	return fizzBuzzHandler{
		fizzBuzzService: fizzBuzzService,
		statsService:    statsService,
	}
}

type FizzBuzzRequestParams struct {
	Limit                     int    `json:"limit"  binding:"required,gt=1"`
	FirstMultipleOfToReplace  int    `json:"int1" binding:"required,gt=0"`
	FirstValueToReplaceWith   string `json:"str1" binding:"required"`
	SecondMultipleOfToReplace int    `json:"int2" binding:"required,gt=0"`
	SecondValueToReplaceWith  string `json:"str2" binding:"required"`
}

type StatsResponse struct {
	Request   FizzBuzzRequestParams `json:"topRequest"`
	TotalHits int64                 `json:"hits"`
}

// @Summary Generate 'FizzBuzz' array
// @Description Generate 'FizzBuzz' array
// @id fizzbuzz
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /fizzbuzz [POST]
func (handler fizzBuzzHandler) handleFizzBuzz(ctx *gin.Context) {
	requestParams := FizzBuzzRequestParams{}
	if err := ctx.ShouldBind(&requestParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Notify stats service to update counter for the given request
	request, err := json.Marshal(requestParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	handler.statsService.Hit(request)

	result := handler.fizzBuzzService.FizzBuzz(toFizzBuzzParams(requestParams))
	ctx.JSON(http.StatusOK, result)
}

// @Summary Get top FizzBuzz request and total hits
// @Description Get top FizzBuzz request and total hits
// @id top-requests
// @Accept json
// @Produce json
// @Success 200 {object} StatsResponse
// @Failure 404
// @Failure 500 {object} object
// @Router /stats [POST]
func (handler fizzBuzzHandler) handleTopRequests(ctx *gin.Context) {
	request, hits := handler.statsService.GetTopRequest()

	if request != nil {
		requestParams := FizzBuzzRequestParams{}
		if err := json.Unmarshal(request, &requestParams); err != nil {
			errors.Wrap(err, "failed to parse top request")
			ctx.JSON(http.StatusInternalServerError,
				gin.H{
					"error": errors.Wrap(err, "failed to parse top request").Error(),
				})
			return
		}

		ctx.JSON(http.StatusOK, StatsResponse{
			Request:   requestParams,
			TotalHits: hits,
		})
	}

	ctx.AbortWithStatus(http.StatusNotFound)
}

func toFizzBuzzParams(request FizzBuzzRequestParams) *models.FizzBuzz {
	return &models.FizzBuzz{
		Limit: request.Limit,
		Fizz: models.ReplaceMultipleOfWith{
			MutlipleOf:  request.FirstMultipleOfToReplace,
			ReplaceWith: request.FirstValueToReplaceWith,
		},
		Buzz: models.ReplaceMultipleOfWith{
			MutlipleOf:  request.SecondMultipleOfToReplace,
			ReplaceWith: request.SecondValueToReplaceWith,
		},
	}
}
