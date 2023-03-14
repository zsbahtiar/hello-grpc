package entity

type ValidateTokenRequest struct {
	Token string
}

type ValidationTokenResponse struct {
	Name  string
	Email string
}
