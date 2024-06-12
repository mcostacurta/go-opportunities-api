package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcostacurta/gooportunities/schemas"
)

// @BasePath /api/v1

// @summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening Identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/{id} [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}

	openning := schemas.Opening{}

	if err := db.First(&openning, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	if err := db.Delete(&openning).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", openning)
}
