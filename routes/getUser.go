package routes

import (
	"database/sql"
	"os"

	"github.com/billzayy/crud-go/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetAllUser() ([]models.User, error) {
	godotenv.Load()

	DBSrc, err := sql.Open("mysql", os.Getenv("DBPATH"))

	if err != nil {
		return nil, err
	}

	db, err := DBSrc.Query("SELECT * FROM User")

	if err != nil {
		return nil, err
	}

	var userList []models.User

	for db.Next() {
		var list models.User

		err := db.Scan(&list.Id, &list.UserName, &list.Password, &list.Email)

		if err != nil {
			return nil, err
		}

		userList = append(userList, list)
	}

	return userList, nil
}

func GetUserById(id int) (models.User, error) {
	godotenv.Load()

	DBSrc, err := sql.Open("mysql", os.Getenv("DBPATH"))

	if err != nil {
		return models.User{}, err
	}

	db, err := DBSrc.Query("SELECT * FROM User Where Id = ?", id)

	if err != nil {
		return models.User{}, err
	}

	var userList models.User

	for db.Next() {
		var list models.User

		err := db.Scan(&list.Id, &list.UserName, &list.Password, &list.Email)

		if err != nil {
			return models.User{}, err
		}

		userList = list
	}

	return userList, nil
}
