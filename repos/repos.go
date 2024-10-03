package repos

import (
	"apartmentRent/models"

	"gorm.io/gorm"
)

type Repos struct {
	db *gorm.DB
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		db: db,
	}
}

// Function located in repo, used by services
type IRepos interface {
	// Users - Main functions
	CreateUser(models.Users) error
	GetUsers() ([]models.Users, error)

	// Users - Support functions
	EmailExists(string) bool
	GetUserByField(string, any) (models.Users, error)
	AllowAccess(int, string) bool

	// Daily rents - Main functions
	CreateDailyRent(models.DailyRents) error
	GetDailyRents() ([]models.DailyRents, error)
	GetDailyRent(int) (models.DailyRents, error)
	UpdateDailyRent(models.DailyRents) error
	DeleteDailyRent(int) error

	// Daily rents - Support functions
	DailyRentExists(string) bool
	GetDailyRentByField(string, any) (models.DailyRents, error)

	// Long term rents - Main functions
	CreateLongTermRent(models.LongTermRents) error
	GetLongTermRents() ([]models.LongTermRents, error)
	GetLongTermRent(int) (models.LongTermRents, error)
	UpdateLongTermRent(models.LongTermRents) error
	DeleteLongTermRent(int) error

	// Long term rents - Support functions
	LongTermRentExists(string) bool
	GetLongTermRentByField(string, any) (models.LongTermRents, error)

	// Markets - Main functions
	CreateMarket(models.Markets) error
	GetMarkets() ([]models.Markets, error)
	GetMarket(int) (models.Markets, error)
	UpdateMarket(models.Markets) error
	DeleteMarket(int) error

	// Markets - Support functions
	MarketExists(string) bool
	GetMarketByField(string, any) (models.Markets, error)
}
