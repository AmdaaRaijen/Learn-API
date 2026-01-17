package auth

import "errors"

type registerRequestParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r registerRequestParams) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.Email == "" {
		return errors.New("email is reqired")
	}

	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
