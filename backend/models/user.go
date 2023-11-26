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

		err := rows.Scan()
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}
