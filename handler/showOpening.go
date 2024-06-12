package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcostacurta/gooportunities/schemas"
)

// @BasePath /api/v1

// @summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening Identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings/{id} [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "parameter").Error())
		return
	}

	openning := schemas.Opening{}
	if err := db.Find(&openning, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if openning.ID == 0 {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", openning)
}
