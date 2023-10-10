package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billzayy/crud-go/models"
	"github.com/billzayy/crud-go/routes"
	"github.com/gin-gonic/gin"
)

// UpdateUserController 	godoc
//
//	@Summary		Update User
//	@Description	Update user
//	@Tags			User accounts
//	@Accept			json
//	@Produce		json
//	@Param			userInfo	body		models.User	true	"Update User"
//	@Success		200			{object}	models.User
//	@Router			/update-user [put]
func UpdateUserController(c *gin.Context) {
	var update models.User

	err := json.NewDecoder(c.Request.Body).Decode(&update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failure", "data": err})
	}

	updateResult, err := routes.UpdateUser(update)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failure", "data": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": updateResult})
	}
}
