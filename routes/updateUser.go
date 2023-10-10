package routes

import (
	"database/sql"
	"os"

	"github.com/billzayy/crud-go/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func UpdateUser(input models.User) (string, error) {
	godotenv.Load()

	dbSrc, err := sql.Open("mysql", os.Getenv("DBPATH"))

	if err != nil {
		return "", err
	}

	dbPrep, err := dbSrc.Prepare("UPDATE User SET UserName = ?, Password = ?, Email = ? WHERE Id = ?")

	if err != nil {
		return "", err
	}

	_, err = dbPrep.Exec(input.UserName, input.Password, input.Email, input.Id)

	if err != nil {
		return "", err
	}

	return "Updated Successful", nil
}
