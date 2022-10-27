package controller

import (
	"backendmod/entity"
	"backendmod/types"
	"encoding/json"
	"net/http"
	"time"
)

func GoogleAuthenticator(w http.ResponseWriter, r *http.Request) {
	var callBackData types.GoogleCallbackType

	err := json.NewDecoder(r.Body).Decode(&callBackData)
	if err != nil {
		failedGoogleAuthResponse(w, "Invalid format.")
		return
	}

	user, userAlreadyExists := userModel.GetByEmail(callBackData.Email)

	if user.Google == "1" {
		tokenString, _ := CreateToken(w, callBackData.Email)
		sucessfullGoogleAuthResponse(w, user, callBackData, tokenString)
		return
	}

	if userAlreadyExists && user.Google == "0" {
		_, userDataUpdated := googleUser.SET(callBackData.Email, "1")

		if userDataUpdated {
			tokenString, _ := CreateToken(w, callBackData.Email)
			sucessfullGoogleAuthResponse(w, user, callBackData, tokenString)
			return
		} else {
			failedGoogleAuthResponse(w, "Couldn't authenticate with Google account.")
			return
		}
	}

	if !userAlreadyExists {
		_, createdGoogleAccount := googleUser.CREATE(callBackData.Email, callBackData.Name)
		if createdGoogleAccount {
			tokenString, _ := CreateToken(w, callBackData.Email)
			sucessfullGoogleAuthResponse(w, user, callBackData, tokenString)
			return
		} else {
			failedGoogleAuthResponse(w, "Couldn't authenticate with Google account.")
			return
		}
	}

	failedGoogleAuthResponse(w, "Couldn't authenticate with Google account.")
	return
}

func sucessfullGoogleAuthResponse(w http.ResponseWriter, user entity.User, callback types.GoogleCallbackType, tokenString string) {
	json.NewEncoder(w).Encode(types.SuccessGoogleAuthResponseType{
		Success: true,
		User: types.GoogleUserType{
			Email:     callback.Email,
			Full_Name: user.Full_Name,
			Telephone: user.Telephone,
			Google:    "1",
		},
		Token:     tokenString,
		ExpiresAt: time.Now().Add(time.Minute * 30),
		Message:   "User has been successfully authenticated with Google account.",
	})
}

func failedGoogleAuthResponse(w http.ResponseWriter, str string) {
	json.NewEncoder(w).Encode(types.FailedGoogleAuthResponseType{
		Success: false,
		Message: str,
	})
}
