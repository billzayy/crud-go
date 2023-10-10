package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/billzayy/crud-go/routes"
	"github.com/gin-gonic/gin"
)

// GetUserController 	godoc
//
//	@Summary		Show All User
//	@Description	get all user
//	@Tags			User accounts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Router			/get-all-user [get]
func GetUserController(c *gin.Context) {
	getAllUser, err := routes.GetAllUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": getAllUser})
	}
}

// GetUserById 	godoc
//
//	@Summary		Show User By Id
//	@Description	get user by id
//	@Tags			User accounts
//	@Param			id	path	int	true	"Account ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Router			/get-user/{id} [get]
func GetUserById(c *gin.Context) {
	inputParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	}

	result, err := routes.GetUserById(inputParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": result})
	}
}
