package controllers

import (
	"echo-rest-api/lib"
	"echo-rest-api/models"
	"echo-rest-api/storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Lists all users
func GetUsers(c echo.Context) error {

	// Parsing limit parameter
	limit, err := lib.QueryParamToInt(c.QueryParam("limit"), 50)
	if err != nil {
		errorMessage := lib.Message{Message: "Failed to parse limit " + c.QueryParam("limit")}
		return c.JSON(http.StatusBadRequest, errorMessage)
	}
	switch {
	case limit < 1:
		limit = 1
	case limit > 100:
		limit = 100
	}

	// Parsing page parameter
	page, err := lib.QueryParamToInt(c.QueryParam("page"), 1)
	if err != nil {
		errorMessage := lib.Message{Message: "Failed to parse page " + c.QueryParam("page")}
		return c.JSON(http.StatusBadRequest, errorMessage)
	}
	switch {
	case page < 1:
		page = 1
	}

	// Paginating database
	db := storage.GetDBInstance()
	scope := lib.Paginate(page, limit)

	// Querying
	users := []models.User{}
	if err := db.Scopes(scope).Find(&users).Error; err != nil {
		errorMessage := lib.Message{Message: "Failed to query database: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	return c.JSON(http.StatusOK, users)
}

// Getting a single user
func GetSingleUser(id int) (models.User, error) {
	db := storage.GetDBInstance()
	user := models.User{}

	err := db.First(&user, id).Error
	return user, err
}

// Get a single user
func GetUser(c echo.Context) error {
	// Parsing ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := lib.Message{Message: "Failed parsing id " + c.Param("id")}
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	// Querying
	user, err := GetSingleUser(id)
	if err != nil {
		errorMessage := lib.Message{Message: "Failed to query database: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	return c.JSON(http.StatusOK, user)
}

// Create
func CreateUser(c echo.Context) error {
	// Parse input data
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		errorMessage := lib.Message{Message: "Failed to parse new user data: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	fmt.Println(user) // debug

	// Creating new entry
	db := storage.GetDBInstance()
	if err := db.Create(&user).Error; err != nil {
		errorMessage := lib.Message{Message: "Failed to create new user: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}
	return c.JSON(http.StatusCreated, user)
}

// Update
func UpdateUser(c echo.Context) error {
	// Parsing ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := lib.Message{Message: "Failed parsing id " + c.Param("id")}
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	// Querying
	user, err := GetSingleUser(id)
	if err != nil {
		errorMessage := lib.Message{Message: "Failed to query database: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	// Parse input data
	if err := c.Bind(&user); err != nil {
		errorMessage := lib.Message{Message: "Failed to parse new user data: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	// Updating entry
	db := storage.GetDBInstance()
	if err := db.Updates(user).Error; err != nil {
		errorMessage := lib.Message{Message: "Failed to update user: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}
	return c.JSON(http.StatusCreated, user)
}

// Delete
func DeleteUser(c echo.Context) error {
	// Parsing ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := lib.Message{Message: "Failed parsing id " + c.Param("id")}
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	// Querying
	user, err := GetSingleUser(id)
	if err != nil {
		errorMessage := lib.Message{Message: "Failed to query database: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	// Deleting entry
	db := storage.GetDBInstance()
	db.Delete(&user)
	return c.NoContent(http.StatusNoContent)
}
