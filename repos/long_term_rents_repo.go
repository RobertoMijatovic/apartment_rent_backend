package repos

import (
	"apartmentRent/helpers"
	"apartmentRent/models"
	"fmt"
)

// Main functions
func (r *Repos) CreateLongTermRent(input models.LongTermRents) error {
	if err := r.db.Create(&input).Error; err != nil {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) GetLongTermRents() ([]models.LongTermRents, error) {
	var longTermRents []models.LongTermRents

	if res := r.db.Find(&longTermRents); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return longTermRents, nil
}

func (r *Repos) GetLongTermRent(id int) (models.LongTermRents, error) {
	var longTermRent models.LongTermRents

	if err := r.db.Where("id = ?", id).First(&longTermRent).Error; err != nil {
		return longTermRent, helpers.ErrRepo
	}

	return longTermRent, nil
}

func (r *Repos) UpdateLongTermRent(input models.LongTermRents) error {
	res := r.db.Model(&models.LongTermRents{}).Where("id = ?", input.ID).Updates(&input)
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) DeleteLongTermRent(id int) error {
	res := r.db.Where("id = ?", id).Delete(&models.LongTermRents{})
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

// Support functions
func (r *Repos) LongTermRentExists(name string) bool {
	res := r.db.Where("name = ?", name).First(&models.LongTermRents{}).RowsAffected

	return res != 0
}

func (r *Repos) GetLongTermRentByField(field string, value any) (models.LongTermRents, error) {
	var longTermRents models.LongTermRents

	if err := r.db.Where(fmt.Sprintf("%v = ?", field), value).First(&longTermRents).Error; err != nil {
		return longTermRents, helpers.ErrRepo
	}

	return longTermRents, nil
}
