package forms

type User struct {
	ID uint64 `json:"id"`
	Username string `json:"username"`
	DisplayName string `json:"display_name"`
	CreatedDate string `json:"created_at"`
}
