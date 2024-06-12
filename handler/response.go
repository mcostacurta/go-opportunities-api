package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcostacurta/gooportunities/schemas"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

func sendCreateSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json: "message"`
	Data    schemas.OpeningResponse `json: "data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json: "message"`
	Data    schemas.OpeningResponse `json: "data"`
}

type ShowOpeningResponse struct {
	Message string                  `json: "message"`
	Data    schemas.OpeningResponse `json: "data"`
}

type ListOpeningsResponse struct {
	Message string                    `json: "message"`
	Data    []schemas.OpeningResponse `json: "data"`
}

type UpdateOpeningResponse struct {
	Message string                  `json: "message"`
	Data    schemas.OpeningResponse `json: "data"`
}
