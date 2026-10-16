package validator

import (
	"github.com/griittyy/forum/internal/models"
)

type RegValidationErrors struct {
	Email           string `json:"email,omitzero"`
	Username        string `json:"username,omitzero"`
	Name            string `json:"name,omitzero"`
	Password        string `json:"password,omitzero"`
	ConfirmPassword string `json:"confirm-password,omitzero"`
}

type RegValidationResult struct {
	Message string              `json:"message"`
	Errors  RegValidationErrors `json:"errors"`
}

func ValidateRegistrationData(userRegData models.UserRegistrationData) RegValidationResult {
	var errors RegValidationErrors

	if IsEmpty(userRegData.Email) {
		errors.Email = "This field cannot be empty"
	} else if !IsValidEmail(userRegData.Email) {
		errors.Email = "The provided email address is not valid"
	}

	if IsEmpty(userRegData.Name) {
		errors.Name = "This field cannot be empty"
	} else if !HasMinLength(userRegData.Name, 4) {
		errors.Name = "The length of the provided name must be four characters or more."
	}

	if IsEmpty(userRegData.Username) {
		errors.Username = "This field cannot be empty"
	} else if !HasMinLength(userRegData.Username, 3) {
		errors.Username = "The length of the provided username must be 3 characters or more."
	}

	if IsEmpty(userRegData.Password) {
		errors.Password = "This field cannot be empty"
	} else if ContainsSpaces(userRegData.Password) {
		errors.Password = "This field cannot contain spaces"
	} else if !HasMinLength(userRegData.Password, 8) {
		errors.Password = "The length of the provided password must be 8 characters or more."
	}

	if IsEmpty(userRegData.ConfirmPassword) {
		errors.ConfirmPassword = "This field cannot be empty"
	} else if !Matches(userRegData.Password, userRegData.ConfirmPassword) {
		errors.ConfirmPassword = "This field must match the provided password"
	}

	return RegValidationResult{
		Message: "There were some errors with your provided json",
		Errors:  errors,
	}
}

func (userRegdata RegValidationResult) Valid() bool {
	return IsEmptyStruct(userRegdata)
}
