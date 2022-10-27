package controller

import (
	"backendmod/types"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	form, notValidMessage, formIsNotValid := validateLoginForm(w, r)
	if formIsNotValid {
		json.NewEncoder(w).Encode(types.LoginResponseType{
			Success: formIsNotValid,
			Message: notValidMessage,
		})
		return
	}

	user, userExists := userModel.GetByEmail(form.Email)
	if !userExists {
		json.NewEncoder(w).Encode(types.LoginResponseType{
			Success: false,
			Message: "username entered does not exist.",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		json.NewEncoder(w).Encode(types.LoginResponseType{
			Success: false,
			Message: "password is incorrect.",
		})
		return
	}

	token, err := CreateToken(w, user.Email)
	if err != nil {
		json.NewEncoder(w).Encode(types.LoginResponseType{
			Success: false,
			Message: "Couldn't create token.",
		})
		return
	}

	json.NewEncoder(w).Encode(types.LoginSuccesfulResponseType{
		Success: true,
		User: types.LoginUserType{
			ID:        user.ID_User,
			Email:     user.Email,
			Full_Name: user.Full_Name,
			Telephone: user.Telephone,
			Google:    user.Google,
		},
		Token:   token,
		Message: "User has been successfully authenticated.",
	})
	return
}

func validateLoginForm(w http.ResponseWriter, r *http.Request) (types.LoginFormType, string, bool) {
	var form types.LoginFormType

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		return form, "Invalid login format.", true
	}

	if form.Email == "" || form.Password == "" {
		return form, "Please enter a valid email and password.", true
	}

	return form, "", false
}
