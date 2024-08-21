package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/handlers"
	"github.com/keliMuthengi/invoiving-api/repo/models"
)

func DoGetUsers(c *gin.Context) {
	// get query params;
	pagestr := c.Query("page")
	limitstr := c.Query("limit")
	searchValue := c.Query("searchValue")
	page, err := strconv.Atoi(pagestr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitstr)
	if err != nil || limit < 1 {
		limit = 10
	}

	// calculate the offset;
	page = (page - 1) * limit

	var params = handlers.RequestParams{
		Page:        page,
		Limit:       limit,
		SearchValue: searchValue,
	}

	users, err := handlers.ListUsers(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var userRequest handlers.CreateUserRequest

	// bind json;

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	// pass data to the create user handler;
	userData := models.User{
		Username:    userRequest.Username,
		Email:       userRequest.Email,
		Address:     userRequest.Address,
		Phonenumber: userRequest.Phonenumber,
		Password:    userRequest.Password,
	}

	user, err := handlers.DoCreateUser(userData)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	c.JSON(http.StatusCreated, handlers.ResponseHandler{Message: "User Created Successfully", Data: user})
}

func LoginUser(c *gin.Context) {

	var loginrequest handlers.LoginHandler
	var user models.User
	// bind json with required tags;

	if err := c.ShouldBindJSON(&loginrequest); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	// find user with provided details;
	if usererr := database.DB.Find(&user, "email = ?", loginrequest.Email).Error; usererr != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: usererr.Error(), Status: 1})
		return
	}
	user.LPassword = loginrequest.Password
	token, loginerr := handlers.DoLoginUser(user)

	if loginerr != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: loginerr.Error(), Status: 1})
		return
	}

	c.JSON(http.StatusOK, handlers.ResponseHandler{Message: "User Logged In Successfully", Status: 0, Token: token})
}

func GetTenants(c *gin.Context) {

	// should bind json

	tenants, err := handlers.DoGetTenants()

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	c.JSON(http.StatusOK, handlers.ResponseHandler{Message: "Tenants found!", Status: 0, Data: tenants})
}
