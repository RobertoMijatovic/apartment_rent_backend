package services

import (
	"apartmentRent/helpers"
	"apartmentRent/models"

	"github.com/gin-gonic/gin"
)

func (s *Services) CreateMarket(ctx *gin.Context, input models.Markets) error {
	userID := ctx.Keys["userID"].(int)

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
		return helpers.ErrService
	}

	if marketEx := s.Repos.MarketExists(input.Name); marketEx {
		return helpers.ErrService
	}

	return s.Repos.CreateMarket(input)
}

func (s *Services) GetMarkets(ctx *gin.Context) ([]models.Markets, error) {
	return s.Repos.GetMarkets()
}

func (s *Services) GetMarket(ctx *gin.Context, id int) (models.Markets, error) {
	return s.Repos.GetMarket(id)
}

func (s *Services) UpdateMarket(ctx *gin.Context, input models.Markets) error {
	userID := ctx.Keys["userID"].(int)

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
		return helpers.ErrService
	}

	return s.Repos.UpdateMarket(input)
}

func (s *Services) DeleteMarket(ctx *gin.Context, id int) error {
	userID := ctx.Keys["userID"].(int)

	if !s.Repos.AllowAccess(userID, "admin") {
		return helpers.ErrService
	}

	return s.Repos.DeleteMarket(id)
}
