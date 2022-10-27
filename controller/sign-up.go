package controller

import (
	"backendmod/types"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	var registerForm types.RegisterFormType

	err := json.NewDecoder(r.Body).Decode(&registerForm)
	if err != nil {
		log.Println("Error decoder:", err.Error())
		json.NewEncoder(w).Encode(types.SignUpFailedResponseType{
			Success: false,
			Message: "Invalid format.",
		})
		return
	}

	email := registerForm.Email
	password := registerForm.Password

	if email == "" || password == "" {
		json.NewEncoder(w).Encode(types.SignUpFailedResponseType{
			Success: false,
			Message: "Please enter a valid email and password.",
		})
		return
	}

	_, userAlreadyExists := userModel.GetByEmail(email)

	if userAlreadyExists {
		json.NewEncoder(w).Encode(types.SignUpFailedResponseType{
			Success: false,
			Message: "User already exists.",
		})
		return
	}

	encriptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	_, userCreated := userModel.Create(email, encriptedPassword)

	if userCreated {
		tokenString, err := CreateToken(w, email)
		if err != nil {
			json.NewEncoder(w).Encode(types.SignUpFailedResponseType{
				Success: false,
				Message: "Couldn't create token.",
			})
			return
		} else {
			json.NewEncoder(w).Encode(types.SignUpSuccesfulResponseType{
				Success: true,
				User: types.SignUpUserCreatedType{
					Email:  email,
					Google: "0",
				},
				Token:     tokenString,
				ExpiresAt: time.Now().Add(time.Minute * 30),
				Message:   "User has been successfully created.",
			})
			return
		}
	}

	json.NewEncoder(w).Encode(types.SignUpFailedResponseType{
		Success: false,
		Message: "Could not create user.",
	})
	return
}
