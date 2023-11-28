package forms

type History struct {
	HistoryID uint64 `json:"history_id"`
	QuizID uint64 `json:"quiz_id"`
	CategoryID uint64 `json:"category_id"`
	UserID uint64 `json:"user_id"`
	CreatorID uint64 `json:"user_creator_id"`
	Title string `json:"title"`
	Score uint `json:"score"`
}
