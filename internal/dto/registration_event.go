package dto

type UserRegisteredEvent struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Time     string `json:"time,omitempty"`
}
