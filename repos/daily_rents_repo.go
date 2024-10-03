package repos

import (
	"apartmentRent/helpers"
	"apartmentRent/models"
	"fmt"
	"log"
)

// Main functions
func (r *Repos) CreateDailyRent(input models.DailyRents) error {
	if err := r.db.Create(&input).Error; err != nil {
		log.Printf("Error creating daily rent: %v", err) // Log the actual error
		return helpers.ErrRepo
	}
	return nil
}

func (r *Repos) GetDailyRents() ([]models.DailyRents, error) {
	var dailyRents []models.DailyRents

	if res := r.db.Find(&dailyRents); res.Error != nil {
		log.Printf("Error fetching daily rents: %v", res.Error) // Log the actual error
		return nil, helpers.ErrRepo
	}
	return dailyRents, nil
}

func (r *Repos) GetDailyRent(id int) (models.DailyRents, error) {
	var dailyRent models.DailyRents

	if err := r.db.Where("id = ?", id).First(&dailyRent).Error; err != nil {
		log.Printf("Error fetching daily rent with ID %d: %v", id, err) // Log the actual error
		return dailyRent, helpers.ErrRepo
	}
	return dailyRent, nil
}

func (r *Repos) UpdateDailyRent(input models.DailyRents) error {
	res := r.db.Model(&models.DailyRents{}).Where("id = ?", input.ID).Updates(&input)
	if res.Error != nil || res.RowsAffected == 0 {
		log.Printf("Error updating daily rent ID %d: %v", input.ID, res.Error) // Log the actual error
		return helpers.ErrRepo
	}
	return nil
}

func (r *Repos) DeleteDailyRent(id int) error {
	res := r.db.Where("id = ?", id).Delete(&models.DailyRents{})
	if res.Error != nil || res.RowsAffected == 0 {
		log.Printf("Error deleting daily rent ID %d: %v", id, res.Error) // Log the actual error
		return helpers.ErrRepo
	}
	return nil
}

// Support functions
func (r *Repos) GetDailyRentByField(field string, value any) (models.DailyRents, error) {
	var dailyRent models.DailyRents

	if err := r.db.Where(fmt.Sprintf("%v = ?", field), value).First(&dailyRent).Error; err != nil {
		log.Printf("Error fetching daily rent by field %s: %v", field, err) // Log the actual error
		return dailyRent, helpers.ErrRepo
	}
	return dailyRent, nil
}

func (r *Repos) DailyRentExists(name string) bool {
	res := r.db.Where("name = ?", name).First(&models.DailyRents{}).RowsAffected
	return res != 0
}
