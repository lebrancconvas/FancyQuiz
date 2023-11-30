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

func (q Quiz) GetAllQuizHeader() ([]forms.QuizHeader, error) {
	db := db.GetDB()

	var quizHeaders []forms.QuizHeader

	stmt, err := db.Prepare(`
		SELECT DISTINCT fk_user_id, quiz_id, fk_quiz_category_id, title, description
		FROM quizzes
		WHERE used_flg = true
		ORDER BY created_at DESC
	`)
	if err != nil {
		return quizHeaders, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return quizHeaders, err
	}
	defer rows.Close()

	for rows.Next() {
		var quizHeader forms.QuizHeader
		err := rows.Scan(&quizHeader.UserID, &quizHeader.QuizID, &quizHeader.QuizCategoryID, &quizHeader.Title, &quizHeader.Description)
		if err != nil {
			return quizHeaders, err
		}
		quizHeaders = append(quizHeaders, quizHeader)
	}

	return quizHeaders, nil
}

func (q Quiz) GetAllQuiz() ([]forms.Quiz, error) {
	db := db.GetDB()

	var quizzes []forms.Quiz
	var questions []forms.Question
	var choices []forms.Choice

	stmt, err := db.Prepare(`
		SELECT DISTINCT fk_user_id, quiz_id, fk_quiz_category_id, title, description
		FROM quizzes
		WHERE used_flg = true
	`)
	if err != nil {
		return quizzes, err
	}
	defer stmt.Close()

	questionStmt, err := db.Prepare(`
		SELECT DISTINCT quiz_question_id, question
		FROM quiz_questions
		WHERE fk_quiz_id = $1 AND used_flg = true
	`)
	if err != nil {
		return quizzes, err
	}
	defer questionStmt.Close()

	choiceStmt, err := db.Prepare(`
		SELECT DISTINCT quiz_question_choice_id, question, is_correct
		FROM quiz_question_choices
		WHERE fk_quiz_question_id = $1 AND used_flg = true
	`)
	if err != nil {
		return quizzes, err
	}
	defer choiceStmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return quizzes, err
	}
	defer rows.Close()

	for rows.Next() {
		var quiz forms.Quiz
		err := rows.Scan(&quiz.UserID, &quiz.QuizID, &quiz.QuizCategoryID, &quiz.Title, &quiz.Description)
		if err != nil {
			return quizzes, err
		}

		questionRows, err := questionStmt.Query(quiz.QuizID)
		if err != nil {
			return quizzes, err
		}
		defer questionRows.Close()

		for questionRows.Next() {
			var question forms.Question
			err := rows.Scan(&question.QuestionID, &question.Question)
			if err != nil {
				return quizzes, err
			}

			choiceRows, err := choiceStmt.Query(question.QuestionID)
			if err != nil {
				return quizzes, err
			}
			choiceRows.Close()

			for choiceRows.Next() {
				var choice forms.Choice
				err := rows.Scan(&choice.ChoiceID, &choice.Choice, &choice.IsCorrect)
				if err != nil {
					return quizzes, err
				}
				choices = append(choices, choice)
			}

			question = forms.Question{
				QuestionID: question.QuestionID,
				Question: question.Question,
				Choices: choices,
			}

			questions = append(questions, question)
		}

		quiz = forms.Quiz{
			UserID: quiz.UserID,
			QuizID: quiz.QuizID,
			QuizCategoryID: quiz.QuizCategoryID,
			Title: quiz.Title,
			Description: quiz.Description,
			Questions: questions,
		}

		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func (q Quiz) GetQuizByID(quizID uint64) (forms.Quiz, error) {
	db := db.GetDB()

	var quiz forms.Quiz
	var questions []forms.Question
	var choices []forms.Choice

	stmt, err := db.Prepare(`
		SELECT DISTINCT fk_user_id, quiz_id, fk_quiz_category_id, title, description
		FROM quizzes
		WHERE quiz_id = $1 AND used_flg = true
	`)
	if err != nil {
		return quiz, err
	}
	defer stmt.Close()

	questionStmt, err := db.Prepare(`
		SELECT DISTINCT quiz_question_id, question
		FROM quiz_questions
		WHERE fk_quiz_id = $1 AND used_flg = true
	`)
	if err != nil {
		return quiz, err
	}
	defer questionStmt.Close()

	choiceStmt, err := db.Prepare(`
		SELECT DISTINCT quiz_question_choice_id, question, is_correct
		FROM quiz_question_choices
		WHERE fk_quiz_question_id = $1 AND used_flg = true
	`)
	if err != nil {
		return quiz, err
	}
	defer choiceStmt.Close()

	rows, err := stmt.Query(quizID)
	if err != nil {
		return quiz, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&quiz.UserID, &quiz.QuizID, &quiz.QuizCategoryID, &quiz.Title, &quiz.Description)
		if err != nil {
			return quiz, err
		}

		questionRows, err := questionStmt.Query(quizID)
		if err != nil {
			return quiz, err
		}
		defer questionRows.Close()

		for questionRows.Next() {
			var question forms.Question
			err := rows.Scan(&question.QuestionID, &question.Question)
			if err != nil {
				return quiz, err
			}

			choiceRows, err := choiceStmt.Query(question.QuestionID)
			if err != nil {
				return quiz, err
			}
			choiceRows.Close()

			for choiceRows.Next() {
				var choice forms.Choice
				err := rows.Scan(&choice.ChoiceID, &choice.Choice, &choice.IsCorrect)
				if err != nil {
					return quiz, err
				}
				choices = append(choices, choice)
			}

			question = forms.Question{
				QuestionID: question.QuestionID,
				Question: question.Question,
				Choices: choices,
			}

			questions = append(questions, question)
		}
	}

	quiz = forms.Quiz{
			UserID: quiz.UserID,
			QuizID: quiz.QuizID,
			QuizCategoryID: quiz.QuizCategoryID,
			Title: quiz.Title,
			Description: quiz.Description,
			Questions: questions,
		}

	return quiz, nil
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

func (q Quiz) UpdateQuiz(quizID uint64, categoryID uint64, title string, description string) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		UPDATE quizzes
		SET fk_quiz_category = $2, title = $3, description = $4
		WHERE quiz_id = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(quizID, categoryID, title, description)
	if err != nil {
		return err
	}

	return nil

}
