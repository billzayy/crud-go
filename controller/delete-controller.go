package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billzayy/crud-go/models"
	"github.com/billzayy/crud-go/routes"
	"github.com/gin-gonic/gin"
)

// DeleteUserController 	godoc
//
//	@Summary		Delete User
//	@Description	Delete user
//	@Tags			User accounts
//	@Accept			json
//	@Produce		json
//	@Param			userInfo	body		models.DeleteUser	true	"Delete User"
//	@Success		200			{object}	models.User
//	@Router			/delete-user [delete]
func DeleteUserController(c *gin.Context) {
	var deleteUser models.User

	err := json.NewDecoder(c.Request.Body).Decode(&deleteUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	}

	fmt.Println(deleteUser)
	deleteResult, err := routes.DeleteUser(deleteUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": deleteResult})
	}
}
