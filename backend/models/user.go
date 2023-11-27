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

		err := rows.Scan(&user.UserID, &user.Username, &user.DisplayName, &user.CreatedDate)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u User) CreateUser(user forms.UserRegister) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		INSERT INTO users (username, display_name)
		VALUES ($1, $2)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.DisplayName)
	if err != nil {
		return err
	}

	return nil
}

func (u User) UpdateUser(userID uint64, displayName string) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		UPDATE users
		SET display_name = $2
		WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, displayName)
	if err != nil {
		return err
	}

	return nil
}

func (u User) DeleteUser(userID uint64) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		UPDATE users
		SET used_flg = false
		WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	return nil
}
