package types

import "time"

type GoogleCallbackType struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Verified_Email bool   `json:"verified_email"`
	Name           string `json:"name"`
	Given_Name     string `json:"given_name"`
	Family_Name    string `json:"family_name"`
	Picture        string `json:"picture"`
	Locale         string `json:"locale"`
}

type FailedGoogleAuthResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GoogleUserType struct {
	Email     string `json:"email"`
	Full_Name string `json:"full_name"`
	Telephone string `json:"telephone"`
	Google    string `json:"google"`
}

type SuccessGoogleAuthResponseType struct {
	Success   bool           `json:"success"`
	User      GoogleUserType `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt time.Time      `json:"expires_at"`
	Message   string         `json:"message"`
}
