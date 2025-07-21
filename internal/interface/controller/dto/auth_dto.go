package dto

type RegisterRequest struct {
	Email    string `json:"email" vaildate:"email , required"`
	Password string `json:"password" vaildate:"min=6 , required"`
}
type LoginRequest struct {
	Email    string `json:"email" vaildate:"email , required"`
	Password string `json:"password" vaildate:"min=6 , required"`
}
