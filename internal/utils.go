package internal

type JWWTPAYLOAD struct {
	USerID string `json:"user_id"`
	Email string `json:"email"`
}

func GenerateJWToken(UserID, Email string) (string, error){
	return "mocked_jwt_token", nil
}