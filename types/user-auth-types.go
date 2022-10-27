package types

type AuthUserResponseType struct {
	Success bool         `json:"id"`
	User    AuthUserType `json:"user"`
}

type AuthUserType struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Full_Name string `json:"full_name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Google    string `json:"google"`
}

type AuthUserFailedResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
