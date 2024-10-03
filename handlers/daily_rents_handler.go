package handlers

import (
	"apartmentRent/helpers"
	"apartmentRent/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateDailyRent(ctx *gin.Context) {
	var newDailyRent models.DailyRents

	if err := ctx.ShouldBind(&newDailyRent); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.CreateDailyRent(ctx, newDailyRent)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"response": "success"})
}

func (h *Handlers) GetDailyRents(ctx *gin.Context) {
	res, err := h.Services.GetDailyRents(ctx)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) GetDailyRent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	res, err := h.Services.GetDailyRent(ctx, id)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) UpdateDailyRent(ctx *gin.Context) {
	var upDailyRent models.DailyRents

	if err := ctx.ShouldBind(&upDailyRent); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.UpdateDailyRent(ctx, upDailyRent)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}

func (h *Handlers) DeleteDailyRent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err = h.Services.DeleteDailyRent(ctx, id)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}
