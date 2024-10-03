package repos

import (
	"apartmentRent/helpers"
	"apartmentRent/models"
	"fmt"
)

// Main functions
func (r *Repos) CreateMarket(input models.Markets) error {
	if err := r.db.Create(&input).Error; err != nil {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) GetMarkets() ([]models.Markets, error) {
	var markets []models.Markets

	if res := r.db.Find(&markets); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return markets, nil
}

func (r *Repos) GetMarket(id int) (models.Markets, error) {
	var market models.Markets

	if err := r.db.Where("id = ?", id).First(&market).Error; err != nil {
		return market, helpers.ErrRepo
	}

	return market, nil
}

func (r *Repos) UpdateMarket(input models.Markets) error {
	res := r.db.Model(&models.Markets{}).Where("id = ?", input.ID).Updates(&input)
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) DeleteMarket(id int) error {
	res := r.db.Where("id = ?", id).Delete(&models.Markets{})
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

// Support functions
func (r *Repos) MarketExists(name string) bool {
	res := r.db.Where("name = ?", name).First(&models.Markets{}).RowsAffected

	return res != 0
}

func (r *Repos) GetMarketByField(field string, value any) (models.Markets, error) {
	var markets models.Markets

	if err := r.db.Where(fmt.Sprintf("%v = ?", field), value).First(&markets).Error; err != nil {
		return markets, helpers.ErrRepo
	}

	return markets, nil
}
