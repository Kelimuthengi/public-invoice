package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/keliMuthengi/invoiving-api/handlers"
	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/repo/models"
)

func AddParent(c *gin.Context) {
	var parentInput ParentInput
	var houseunit models.HouseUnitTypes
	// Bind JSON datag
	if err := c.ShouldBindJSON(&parentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Start a new transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not start transaction"})
		return
	}

	// Create User
	user := models.User{
		Username:    parentInput.Username,
		Email:       parentInput.Email,
		Address:     parentInput.Address,
		Phonenumber: parentInput.Phonenumber,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// check if provided unit exists;
	if err := database.DB.Find(&houseunit, uint(parentInput.HouseUnitNameID)).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	// Add a new Parent
	parent := models.Parent{
		Username:        parentInput.Username,
		Email:           parentInput.Email,
		Address:         parentInput.Address,
		Phonenumber:     parentInput.Phonenumber,
		UserID:          user.ID,
		HouseUnitTypeID: houseunit.ID,
	}

	if err := tx.Create(&parent).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add students
	for _, studentInput := range parentInput.Students {
		student := models.Student{
			Username:        studentInput.Username,
			AdmissionNumber: studentInput.AdmissionNumber,
			Stream:          studentInput.Stream,
			Boardingstatus:  studentInput.Boardingstatus,
			Hostelname:      studentInput.Hostelname,
			ParentID:        parent.ID,
		}
		if err := tx.Create(&student).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not commit transaction"})
		return
	}

	// Retrieve the parent with students preloaded
	var parentWithStudents models.Parent
	if err := database.DB.Preload("Students").First(&parentWithStudents, parent.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": parentWithStudents})
}

func ListUsers(params RequestParams) ([]models.User, error) {
	var users []models.User
	// // query the database;
	query := database.DB.Offset(params.Page).Limit(params.Limit)

	// append searchValue incase of any;
	searchValue := params.SearchValue
	if searchValue != "" {
		query = query.Where("username = ?", searchValue)
	}
	if err := query.Preload("Products").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func DoCreateUser(u models.User) (models.User, error) {

	var user models.User

	// hash generatedPassword

	hashedPass, err := user.HashUserPassword(u.Password)

	if err != nil {
		return models.User{}, err
	}

	// sign up new Users!
	newUser := models.User{
		Username:    u.Username,
		Password:    string(hashedPass),
		Email:       u.Email,
		Phonenumber: u.Phonenumber,
		Address:     u.Address,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return newUser, err
	}
	// compare provided Passwords;
	// compErr := user.CompareUserPasswords(hashedPass, []byte(u.Password))

	// if compErr != nil {

	// 	return models.User{}, compErr

	// }
	// fmt.Println("we have successfully compared Passwords !")
	return newUser, nil
}

func DoLoginUser(u models.User) (string, error) {
	var user models.User

	err := user.CompareUserPasswords([]byte(u.Password), []byte(u.LPassword))

	if err != nil {
		return "", err
	}
	fmt.Println("user", u)
	tokenDetails := JwtToken{
		Username:    u.Username,
		UserID:      u.ID,
		Email:       u.Email,
		Phonenumber: u.Phonenumber,
		Address:     u.Address,
	}
	token, err := GenerateNewToken(tokenDetails)

	if err != nil {
		return token, err
	}

	return token, nil
}

func DoGetTenants() ([]models.Parent, error) {

	var parent []models.Parent
	if err := database.DB.Model(&models.Parent{}).Preload("HouseUnitTypes").Find(&parent).Error; err != nil {
		return []models.Parent{}, err
	}
	return parent, nil
}
