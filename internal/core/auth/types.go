package auth

import (
	"errors"
	"time"
)

type registerRequestParams struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type loginRequestParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	Message string  `json:"message"`
	Data    userDTO `json:"data"`
}

type userDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber *string   `json:"phone_number,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

func (r registerRequestParams) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.Email == "" {
		return errors.New("email is reqired")
	}

	if r.PhoneNumber == "" {
		return errors.New("phone number is required")
	}

	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func (r loginRequestParams) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}

	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
