package models

import (
	"errors"

	"github.com/fortytw2/abdi"
	"github.com/jmoiron/sqlx"
)

// User model
type User struct {
	ID int64

	Username     string
	Email        string
	PasswordHash string

	Admin         bool `json:"-"`
	PasswordReset bool `json:"-"`
	Confirmed     bool `json:"-"`
}

var (
	errEmailInvalidOrTaken = errors.New("email is invalid or taken")
	errUsernameTaken       = errors.New("username is invalid or taken")
	errLoginFailure        = errors.New("username or password is not valid")
	errAlreadyConfirmed    = errors.New("user already confirmed")
)

// CreateUser creates a new, validated user
func CreateUser(username string, email string, password string, db *sqlx.DB) error {
	hash, err := abdi.Hash(password)
	if err != nil {
		return err
	}

	user := &User{
		Username:      username,
		Email:         email,
		PasswordHash:  hash,
		Admin:         false,
		PasswordReset: false,
		Confirmed:     false,
	}

	_, err = db.NamedExec("INSERT INTO users (username, email, passwordhash, admin, passwordreset, confirmed) VALUES (:username, :email, :passwordhash, :admin, :passwordreset, :confirmed)", user)
	if err != nil {
		return err
	}

	return nil
}

// FindUser a user by username
func FindUser(username string, db *sqlx.DB) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CheckToken checks an authentication token user/expiry/etc against db
func CheckToken(token string) (*User, error) {
	return nil, nil
}

// Save a user to the database, optionally validating - ensures no duplicates
// and updates any existing record with the same ID
// nothing persists to the database but (*User).Save()
func (u *User) Save(validate bool) error {
	return nil
}

// CheckPassword checks a users password against the password hash and returns
// a bool and any errors
func (u *User) CheckPassword(password string) (bool, error) {
	return false, nil
}

// StartPasswordReset sends a password reset email and sets passwordReset to true
func (u *User) StartPasswordReset() error {
	return nil
}

// ConfirmPasswordReset sends a password reset email and sets passwordReset to true
func (u *User) ConfirmPasswordReset(prCode string) error {
	return nil
}

// GenAuthToken generates a signed authentication token for the user
func (u *User) GenAuthToken() (*string, error) {
	return nil, nil
}

// GenConfirmationCode creates a confirmationcode using crypto
func (u *User) GenConfirmationCode() (*string, error) {
	return nil, nil
}

// Confirm the user based on the confirmation code passed
func (u *User) Confirm(cc string) error {
	return nil
}

// splits out validation logic
func validate(username string, email string, password string) error {
	return nil
}
