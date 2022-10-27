package entity

type User struct {
	ID_User   string `json:"id_user"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Full_Name string `json:"full_name"`
	Telephone string `json:"telephone"`
	Google    string `json:"id_google"`
}
