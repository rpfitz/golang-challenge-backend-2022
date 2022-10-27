package types

import "time"

type RegisterFormType struct {
	Email    string
	Password string
}

type SignUpFailedResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SignUpUserCreatedType struct {
	Email  string `json:"email"`
	Google string `json:"google"`
}

type SignUpSuccesfulResponseType struct {
	Success   bool                  `json:"success"`
	User      SignUpUserCreatedType `json:"user"`
	Token     string                `json:"token"`
	ExpiresAt time.Time             `json:"expires_at"`
	Message   string                `json:"message"`
}
