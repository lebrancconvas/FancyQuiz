package models

import (
	"github.com/lebrancconvas/FancyQuiz/db"
	"github.com/lebrancconvas/FancyQuiz/forms"
)

type Quiz struct{}

func (q Quiz) GetAllQuizCategory() ([]forms.QuizCategory, error) {
	db := db.GetDB()

	var quizCategories []forms.QuizCategory

	stmt, err := db.Prepare(`
		SELECT DISTINCT quiz_category_id, category
		FROM quiz_categories
		WHERE used_flg = true
		ORDER BY quiz_category_id ASC
	`)
	if err != nil {
		return quizCategories, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return quizCategories, err
	}
	defer rows.Close()

	for rows.Next() {
		var quizCategory forms.QuizCategory
		err := rows.Scan(&quizCategory.QuizCategoryID, &quizCategory.Category)
		if err != nil {
			return quizCategories, err
		}
		quizCategories = append(quizCategories, quizCategory)
	}

	return quizCategories, nil
}
