package forms

type User struct {
	UserID uint64 `json:"user_id"`
	Username string `json:"username"`
	DisplayName string `json:"display_name"`
	CreatedDate string `json:"created_at"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
}
