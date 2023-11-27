package models

import (
	"github.com/lebrancconvas/FancyQuiz/db"
	"github.com/lebrancconvas/FancyQuiz/forms"
)

type User struct{}

func (u User) GetAllUsers() ([]forms.User, error) {
	var users []forms.User

	db := db.GetDB()

	stmt, err := db.Prepare(`
		SELECT DISTINCT id, username, display_name, created_at
		FROM users
		WHERE used_flg = true
		ORDER BY created_at DESC
	`)
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user forms.User

		err := rows.Scan(&user.ID, &user.Username, &user.DisplayName, &user.CreatedDate)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u User) CreateUser(user forms.UserRegister) (uint64, error) {
	var userID uint64

	db := db.GetDB()

	stmt, err := db.Prepare(`
		INSERT INTO users (username, display_name)
		VALUES ($1, $2)
	`)
	if err != nil {
		return userID, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Username, user.DisplayName)
	if err != nil {
		return userID, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			return userID, err
		}
	}

	return userID, nil
}
