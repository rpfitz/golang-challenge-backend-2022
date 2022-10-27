package controller

import (
	"backendmod/config"
	"backendmod/types"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func EditProfile(w http.ResponseWriter, r *http.Request) {
	success, message, userEmailFromRequest, form, fullNameUpdated, telephoneUpdated, emailUpdated, token := handleEditProfile(w, r)

	if success {
		json.NewEncoder(w).Encode(types.UpdateProfileSuccessType{
			Success:            true,
			Message:            "Profile successfully changed.",
			Email_From_Request: userEmailFromRequest,
			Full_Name:          form.Full_Name,
			Telephone:          form.Telephone,
			Email:              form.Email,
			Full_Name_Updated:  fullNameUpdated,
			Telephone_Updated:  telephoneUpdated,
			Email_Updated:      emailUpdated,
			Token:              token,
		})
		return
	}

	if message == "Unauthorized." {
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(types.UpdateProfileFailedType{
		Success: success,
		Message: message,
	})
	return
}

func handleEditProfile(w http.ResponseWriter, r *http.Request) (bool, string, string, types.UpdateProfileSuccessType, bool, bool, bool, string) {

	var (
		fullNameUpdated  = false
		telephoneUpdated = false
		emailUpdated     = false
	)

	isValid, validationMessage, userEmailFromRequest, form := validateEditProfileRequest(r)
	if !isValid {
		success := false
		message := validationMessage
		return success, message, userEmailFromRequest, form, fullNameUpdated, telephoneUpdated, emailUpdated, ""
	}

	_, userExists := userModel.GetByEmail(form.Email)
	if userExists && (userEmailFromRequest != form.Email) {
		success := false
		message := "Email already exists."
		return success, message, userEmailFromRequest, form, fullNameUpdated, telephoneUpdated, emailUpdated, ""
	}

	_, fullNameUpdated = userModel.SET(userEmailFromRequest, form.Full_Name, "full_name")
	_, telephoneUpdated = userModel.SET(userEmailFromRequest, form.Telephone, "telephone")

	if form.Email != "" {
		_, emailUpdated = userModel.SET(userEmailFromRequest, form.Email, "email")
	}

	if !fullNameUpdated && !telephoneUpdated && !emailUpdated {
		success := false
		message := "Didn't modify anything."
		return success, message, userEmailFromRequest, form, fullNameUpdated, telephoneUpdated, emailUpdated, ""
	}

	if fullNameUpdated || telephoneUpdated || emailUpdated {
		success := true
		message := "Profile successfully changed."

		if emailUpdated {
			newToken, err := CreateToken(w, form.Email)
			if err != nil {
				success := false
				message := "Couldn't create token."
				return success, message, userEmailFromRequest, form, fullNameUpdated, telephoneUpdated, emailUpdated, newToken
			}
			return success, message, userEmailFromRequest, form, fullNameUpdated, telephoneUpdated, emailUpdated, newToken
		}

		return success, message, userEmailFromRequest, form, fullNameUpdated, telephoneUpdated, emailUpdated, ""
	}

	success := false
	message := "Something went wrong."
	return success, message, userEmailFromRequest, form, fullNameUpdated, telephoneUpdated, emailUpdated, ""
}

func validateEditProfileRequest(r *http.Request) (bool, string, string, types.UpdateProfileSuccessType) {
	r.Header.Set("Content-Type", "application/json")

	var form types.UpdateProfileSuccessType

	jwtCookie, _ := GetCookie(r, "Authorization")

	token, err := jwt.ParseWithClaims(jwtCookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})
	if err != nil {
		success := false
		message := "Unauthorized."
		requestEmail := ""
		return success, message, requestEmail, form
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, userExists := userModel.GetByEmail(claims.Issuer)

	if !userExists {
		success := false
		message := "Unauthorized email."
		requestEmail := ""
		return success, message, requestEmail, form
	}

	requestEmail := user.Email

	err1 := json.NewDecoder(r.Body).Decode(&form)
	if err1 != nil {
		return false, "Invalid format.", requestEmail, form
	}

	if form.Email != "" {
		emailFormatMessage, emailFormatIsValid := validateEmailFormat(form.Email)
		if !emailFormatIsValid {
			return emailFormatIsValid, emailFormatMessage, requestEmail, form
		}
	}

	return true, "", requestEmail, form
}

func validateEmailFormat(emailToCheck string) (string, bool) {
	if emailToCheck == "a" {
		return "You must enter a valid email.", false
	}

	return "", true
}
