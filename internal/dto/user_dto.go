package dto

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserMinDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTOutput struct {
	AccessToken string     `json:"access_token"`
	Payload     JWTPayload `json:"payload"`
}
type JWTPayload struct {
	Sub string `json:"sub"`
}
