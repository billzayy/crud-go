package routes

import (
	"database/sql"
	"os"

	"github.com/billzayy/crud-go/models"
	"github.com/joho/godotenv"
)

func CreateUser(input models.User) (string, error) {
	godotenv.Load()

	dbSrc, err := sql.Open("mysql", os.Getenv("DBPATH"))

	if err != nil {
		return "", nil
	}

	dbPrep, err := dbSrc.Prepare("INSERT INTO User (UserName, Password, Email) VALUES (?,?,?)")

	if err != nil {
		return "", nil
	}

	_, err = dbPrep.Exec(input.UserName, input.Password, input.Email)

	if err != nil {
		return "", nil
	}

	return "Sign Up Successful", nil
}
