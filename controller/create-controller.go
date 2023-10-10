package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billzayy/crud-go/models"
	"github.com/billzayy/crud-go/routes"
	"github.com/gin-gonic/gin"
)

// CreateUserController 	godoc
//
//	@Summary		Create User
//	@Description	Create user
//	@Tags			User accounts
//	@Accept			json
//	@Produce		json
//	@Param			userInfo	body		models.AddUser	true	"New User"
//	@Success		200			{object}	models.User
//	@Router			/create-user [post]
func CreateUserController(c *gin.Context) {
	var inputCreate models.User

	err := json.NewDecoder(c.Request.Body).Decode(&inputCreate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failure", "data": err})
	}

	createResult, err := routes.CreateUser(inputCreate)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failure", "data": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": createResult})
	}
}
