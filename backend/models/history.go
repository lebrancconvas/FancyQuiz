package models

import (
	"github.com/lebrancconvas/FancyQuiz/db"
	"github.com/lebrancconvas/FancyQuiz/forms"
)

type History struct{}

func (h History) GetAllHistoryFromUser(userID uint64) ([]forms.History, error) {
	db := db.GetDB()

	var histories []forms.History

	stmt, err := db.Prepare(`
		SELECT quiz_histories.quiz_history_id, quiz_histories.fk_quiz_id, quiz_histories.fk_quiz_category_id, quiz_histories.fk_quiz_participant_user_id, quiz_histories.fk_quiz_creator_user_id, quizzes.quiz_title, quiz_histories.score
		FROM quiz_histories
		INNER JOIN quizzes ON quiz_histories.fk_quiz_id = quizzes.quiz_id
		WHERE quiz_histories.fk_quiz_participant_user_id = $1 AND quiz_histories.used_flg = true
		ORDER BY quiz_histories.created_at DESC
	`)
	if err != nil {
		return histories, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return histories, err
	}
	defer rows.Close()

	for rows.Next() {
		var history forms.History

		err := rows.Scan(&history.HistoryID, &history.QuizID, &history.CategoryID, &history.UserID, &history.CreatorID, &history.Title, &history.Score)
		if err != nil {
			return histories, err
		}

		histories = append(histories, history)
	}

	return histories, nil
}
