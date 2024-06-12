package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcostacurta/gooportunities/schemas"
)

// @BasePath /api/v1

// @summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request Body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	openning := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&openning).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		// return
	}

	sendCreateSuccess(ctx, "create-opening", openning)
}
