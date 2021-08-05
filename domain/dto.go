package domain

type UserRegistrationRequest struct {
	Login string
	Password string
}

type UserRegistrationResponse struct {
	Message string
}