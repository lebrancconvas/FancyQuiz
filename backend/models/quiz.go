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

func (q Quiz) CreateQuiz(userID uint64, categoryID uint64, title string, description string) (uint64, error) {
	db := db.GetDB()

	var quizID uint64

	stmt, err := db.Prepare(`
		INSERT INTO quizzes (fk_user_id, fk_quiz_category_id, title, description)
		VALUES ($1, $2, $3, $4)
	`)
	if err != nil {
		return quizID, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, categoryID, title , description)
	if err != nil {
		return quizID, err
	}

	stmt, err = db.Prepare(`
		SELECT quiz_id
		FROM quizzes
		WHERE fk_user_id = $1 AND fk_quiz_category_id = $2 AND title = $3 AND description = $4
		ORDER BY created_at DESC
		LIMIT 1
	`)
	if err != nil {
		return quizID, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, categoryID, title, description)
	if err != nil {
		return quizID, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&quizID)
		if err != nil {
			return quizID, err
		}
	}

	return quizID, nil
}

func (q Quiz) CreateQuizCategory(category string) (uint64, error) {
	db := db.GetDB()

	var quizCategoryID uint64

	stmt, err := db.Prepare(`
		INSERT INTO quiz_categories (category)
		VALUES ($1)
	`)
	if err != nil {
		return quizCategoryID, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(category)
	if err != nil {
		return quizCategoryID, err
	}

	stmt, err = db.Prepare(`
		SELECT quiz_category_id
		FROM quiz_categories
		WHERE category = $1
		ORDER BY created_at DESC
		LIMIT 1
	`)
	if err != nil {
		return quizCategoryID, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(category)
	if err != nil {
		return quizCategoryID, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&quizCategoryID)
		if err != nil {
			return quizCategoryID, err
		}
	}

	return quizCategoryID, nil
}

func (q Quiz) CreateQuizQuestion(quizID uint64, question string) (uint64, error) {
	db := db.GetDB()

	var quizQuestionID uint64

	stmt, err := db.Prepare(`
		INSERT INTO quiz_questions (fk_quiz_id, question)
		VALUES ($1, $2)
	`)
	if err != nil {
		return quizQuestionID, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(quizID, question)
	if err != nil {
		return quizQuestionID, err
	}

	stmt, err = db.Prepare(`
		SELECT quiz_question_id
		FROM quiz_questions
		WHERE fk_quiz_id = $1 AND question = $2
		ORDER BY created_at DESC
		LIMIT 1
	`)
	if err != nil {
		return quizQuestionID, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(quizID, question)
	if err != nil {
		return quizQuestionID, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&quizQuestionID)
		if err != nil {
			return quizQuestionID, err
		}
	}

	return quizQuestionID, nil
}

func (q Quiz) CreateQuizQuestionChoice(quizQuestionID uint64, choice string, isCorrect bool) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		INSERT INTO quiz_question_choices (fk_quiz_question_id, choice, is_correct)
		VALUES ($1, $2, $3)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(quizQuestionID, choice, isCorrect)
	if err != nil {
		return err
	}

	return nil
}

func (q Quiz) DeleteQuiz(quizID uint64) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		UPDATE quizzes
		SET used_flg = false
		WHERE quiz_id = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(quizID)
	if err != nil {
		return err
	}

	return nil
}
