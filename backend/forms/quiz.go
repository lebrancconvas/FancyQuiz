package forms

type QuizCategory struct {
	QuizCategoryID uint64 `json:"quizCategoryID"`
	Category string `json:"category"`
}

type QuizHeader struct {
	UserID uint64 `json:"user_id"`
	QuizID uint64 `json:"quiz_id"`
	QuizCategoryID uint64 `json:"quiz_category_id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type Quiz struct {
	UserID uint64 `json:"user_id"`
	QuizID uint64 `json:"quiz_id"`
	QuizCategoryID uint64 `json:"quiz_category_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Questions []Question `json:"questions"`
}

type Question struct {
	QuestionID uint64 `json:"question_id"`
	Question string `json:"question"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	ChoiceID uint64 `json:"choice_id"`
	Choice string `json:"choice"`
	IsCorrect bool `json:"is_correct"`
}
