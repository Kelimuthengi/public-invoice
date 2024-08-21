package handlers

import (
	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/repo/models"
)

func DoCreateUnit(hd models.HouseUnitName) (models.HouseUnitName, error) {

	if err := database.DB.Create(&hd).Error; err != nil {

		return hd, err
	}

	return hd, nil
}

func DoCreateHouseUnit(hu models.HouseUnitTypes) (models.HouseUnitTypes, error) {
	if err := database.DB.Create(&hu).Error; err != nil {

		return hu, err
	}

	return hu, nil
}

func Dolistunitsnames() ([]models.HouseUnitName, error) {

	var unitnames []models.HouseUnitName

	if err := database.DB.Find(&unitnames).Error; err != nil {
		return unitnames, nil
	}
	return unitnames, nil
}

func Dolisthousingunits() ([]models.HouseUnitTypes, error) {

	var units []models.HouseUnitTypes

	if err := database.DB.Model(&models.HouseUnitTypes{}).Preload("HouseUnitName").Find(&units).Error; err != nil {
		return units, nil
	}
	return units, nil
}
