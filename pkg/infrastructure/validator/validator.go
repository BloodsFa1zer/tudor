package validator

import (
	"errors"
	"fmt"
	"regexp"

	"gopkg.in/validator.v2"
)

var (
	v = validator.NewValidator()
)

func init() {
	v.SetValidationFunc("email", email)
	v.SetValidationFunc("password", password)
	v.SetValidationFunc("phone", phone)
}

func Validate(s interface{}) error {
	return v.Validate(s)
}

func ValidateString(s string, pattern string) error {
	return v.Valid(s, pattern)
}

func email(v interface{}, param string) error {
	if len(v.(string)) == 0 {
		return nil
	}
	match, err := regexp.MatchString("[a-z0-9]+@[a-z]+\\.[a-z]{2,3}", v.(string))
	if err != nil {
		return fmt.Errorf("something went wrong err: %w. can not pass the email validation", err)
	}
	if !match {
		return errors.New("invalid email address format")
	}
	return nil
}

func password(v interface{}, param string) error {
	str := v.(string)
	if len(str) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(str) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(str) {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(str) {
		return errors.New("password must contain at least one digit")
	}
	if !regexp.MustCompile(`[#?!@$%^&*-]`).MatchString(str) {
		return errors.New("password must contain at least one special character")
	}
	return nil
}

func phone(v interface{}, param string) error {
	if len(v.(string)) == 0 {
		return nil
	}
	match, err := regexp.MatchString("^[\\+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$", v.(string))
	if err != nil {
		return fmt.Errorf("something went wrong err: %w. can not pass the phone number validation", err)
	}
	if !match {
		return errors.New("invalid phone number format")
	}
	return nil
}
