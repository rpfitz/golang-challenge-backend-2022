package types

type LoginFormType struct {
	Email    string
	Password string
}

type LoginResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginSuccesfulResponseType struct {
	Success bool          `json:"success"`
	User    LoginUserType `json:"user"`
	Token   string        `json:"token"`
	Message string        `json:"message"`
}

type LoginUserType struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Full_Name string `json:"full_name"`
	Telephone string `json:"telephone"`
	Google    string `json:"google"`
}
