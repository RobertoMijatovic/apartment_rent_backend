package services

import (
	"apartmentRent/helpers"
	"apartmentRent/models"

	"github.com/gin-gonic/gin"
)

func (s *Services) CreateDailyRent(ctx *gin.Context, input models.DailyRents) error {
	userID := ctx.Keys["userID"].(int)

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
		return helpers.ErrService
	}

	if dailyRentEx := s.Repos.DailyRentExists(input.Name); dailyRentEx {
		return helpers.ErrService
	}

	return s.Repos.CreateDailyRent(input)
}

func (s *Services) GetDailyRents(ctx *gin.Context) ([]models.DailyRents, error) {
	return s.Repos.GetDailyRents()
}

func (s *Services) GetDailyRent(ctx *gin.Context, id int) (models.DailyRents, error) {
	return s.Repos.GetDailyRent(id)
}

func (s *Services) UpdateDailyRent(ctx *gin.Context, input models.DailyRents) error {
	userID := ctx.Keys["userID"].(int)

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
		return helpers.ErrService
	}

	return s.Repos.UpdateDailyRent(input)
}

func (s *Services) DeleteDailyRent(ctx *gin.Context, id int) error {
	userID := ctx.Keys["userID"].(int)

	if !s.Repos.AllowAccess(userID, "admin") {
		return helpers.ErrService
	}

	return s.Repos.DeleteDailyRent(id)
}
