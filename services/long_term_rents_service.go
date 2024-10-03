package services

import (
	"apartmentRent/helpers"
	"apartmentRent/models"

	"github.com/gin-gonic/gin"
)

func (s *Services) CreateLongTermRent(ctx *gin.Context, input models.LongTermRents) error {
	userID := ctx.Keys["userID"].(int)

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
		return helpers.ErrService
	}

	if longTermEx := s.Repos.LongTermRentExists(input.Name); longTermEx {
		return helpers.ErrService
	}

	return s.Repos.CreateLongTermRent(input)
}

func (s *Services) GetLongTermRents(ctx *gin.Context) ([]models.LongTermRents, error) {
	return s.Repos.GetLongTermRents()
}

func (s *Services) GetLongTermRent(ctx *gin.Context, id int) (models.LongTermRents, error) {
	return s.Repos.GetLongTermRent(id)
}

func (s *Services) UpdateLongTermRent(ctx *gin.Context, input models.LongTermRents) error {
	userID := ctx.Keys["userID"].(int)

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
		return helpers.ErrService
	}

	return s.Repos.UpdateLongTermRent(input)
}

func (s *Services) DeleteLongTermRent(ctx *gin.Context, id int) error {
	userID := ctx.Keys["userID"].(int)

	if !s.Repos.AllowAccess(userID, "admin") {
		return helpers.ErrService
	}

	return s.Repos.DeleteLongTermRent(id)
}
