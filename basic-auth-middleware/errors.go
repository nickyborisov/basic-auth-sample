package basic_auth

import "fmt"

type AuthenticationError struct {
}

func (err AuthenticationError) Error() string {
	return "Base authentication error"
}

type InvalidUserError struct {
	AuthenticationError
	User string
}

func (err InvalidUserError) Error() string {
	return fmt.Sprintf("Inalid user name %s", err.User)
}

func NewInvalidUserError(user string) error {
	return InvalidUserError{
		User:user,
	}
}

type InvalidPasswordError struct {
	AuthenticationError
	Password string
}

func (err InvalidPasswordError) Error() string {
	return fmt.Sprintf("Inalid password %s", err.Password)
}

func NewInvalidPasswordError(password string) error {
	return InvalidPasswordError{
		Password:password,
	}
}
