package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/controllers"
	"github.com/keliMuthengi/invoiving-api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(T *testing.T) {

	// handlers.LoadEnvVariable()
	// MOCK TEST GETTING ENV VARIABLES;

	// test creating user function;
	router := gin.Default()
	router.POST("/users", controllers.CreateUser)

	userRequest := handlers.CreateUserRequest{
		Username:    "testuser",
		Email:       "testuser@example.com",
		Address:     "123 Test St",
		Phonenumber: "1234567890",
		Password:    "password",
	}

	// convert payload to json;
	jsonvalue, _ := json.Marshal(userRequest)

	// create a new http Recorder ;
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonvalue))
	req.Header.Set("Content-Type", "application/json")
	// record response using httptest;
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(T, http.StatusCreated, w.Code)

	var response handlers.ResponseHandler
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(T, err)
	assert.Equal(T, "User Created Successfully", response.Message)
	assert.Equal(T, 0, response.Status) // Assuming status 0 means success
}
