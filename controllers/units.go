package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/handlers"
	"github.com/keliMuthengi/invoiving-api/repo/models"
	"gorm.io/gorm"
)

func AddUnitName(c *gin.Context) {
	var unitypeInput handlers.UnitNameInput

	// bind json;

	if err := c.ShouldBindJSON(&unitypeInput); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	unitNameData := models.HouseUnitName{
		UnitTypeName: unitypeInput.Unitname,
	}

	unitdetails, err := handlers.DoCreateUnit(unitNameData)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	c.JSON(http.StatusCreated, handlers.ResponseHandler{Message: "Unit name added successfully", Data: unitdetails})
}

func AddUnits(c *gin.Context) {

	var unitsinput handlers.UnitsInput
	var unitname models.HouseUnitName
	if err := c.ShouldBindJSON(&unitsinput); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	// find existing unit name;

	if err := database.DB.Find(&unitname, unitsinput.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: "Unit Name not found", Status: 1})
			return
		}

		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	unitData := models.HouseUnitTypes{
		Price:           unitsinput.Price,
		HouseNo:         unitsinput.HouseNo,
		HouseUnitNameID: unitname.ID,
	}

	createdUnit, err := handlers.DoCreateHouseUnit(unitData)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: "Unit with provided details already exists", Status: 1})
			return
		}
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	c.JSON(http.StatusAccepted, handlers.ResponseHandler{Message: "Unit created successfully", Data: createdUnit, Status: 0})
}

func Listunitnames(c *gin.Context) {

	houseunitnames, err := handlers.Dolistunitsnames()

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	c.JSON(http.StatusOK, handlers.ResponseHandler{Data: houseunitnames, Message: "House-unit-names found!"})

}

func ListHousingunits(c *gin.Context) {

	housingunits, err := handlers.Dolisthousingunits()

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	c.JSON(http.StatusOK, handlers.ResponseHandler{Data: housingunits, Message: "House-unit-names found!"})
}
