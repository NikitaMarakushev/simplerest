package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func (user *User) BeforeCreate() error {

	if len(user.Password) > 0 {
		enc, err := encryptString(user.Password)
		if err != nil {
			return err
		}

		user.EncryptedPassword = enc
	}

	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
