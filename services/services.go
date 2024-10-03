package services

import (
	"apartmentRent/models"
	"apartmentRent/repos"

	"github.com/gin-gonic/gin"
)

type Services struct {
	Repos repos.IRepos
}

func NewServices(repos repos.IRepos) *Services {
	return &Services{
		Repos: repos,
	}
}

// Function located in services, used by handlers
type IServices interface {
	// Users
	CreateUser(*gin.Context, models.Users) error
	Login(*gin.Context, models.LoginReq) (models.LoginRes, error)
	GetUsers(*gin.Context) ([]models.Users, error)

	// Daily rents
	CreateDailyRent(*gin.Context, models.DailyRents) error
	GetDailyRents(*gin.Context) ([]models.DailyRents, error)
	GetDailyRent(*gin.Context, int) (models.DailyRents, error)
	UpdateDailyRent(*gin.Context, models.DailyRents) error
	DeleteDailyRent(*gin.Context, int) error

	// Long term rents
	CreateLongTermRent(*gin.Context, models.LongTermRents) error
	GetLongTermRents(*gin.Context) ([]models.LongTermRents, error)
	GetLongTermRent(*gin.Context, int) (models.LongTermRents, error)
	UpdateLongTermRent(*gin.Context, models.LongTermRents) error
	DeleteLongTermRent(*gin.Context, int) error

	// Markets
	CreateMarket(*gin.Context, models.Markets) error
	GetMarkets(*gin.Context) ([]models.Markets, error)
	GetMarket(*gin.Context, int) (models.Markets, error)
	UpdateMarket(*gin.Context, models.Markets) error
	DeleteMarket(*gin.Context, int) error
}
