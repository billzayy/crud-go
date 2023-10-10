package routes

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/billzayy/crud-go/models"
	"github.com/joho/godotenv"
)

func DeleteUser(input models.User) (string, error) {
	godotenv.Load()

	dbSrc, err := sql.Open("mysql", os.Getenv("DBPATH"))

	if err != nil {
		return "", err
	}

	dbPrep, err := dbSrc.Prepare("DELETE FROM User WHERE Id = ?")

	if err != nil {
		return "", err
	}

	_, err = dbPrep.Exec(input.Id)
	fmt.Println(input.Id)

	if err != nil {
		return "", err
	}

	return "Delete Successful", nil
}
