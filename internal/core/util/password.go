package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPwd), nil
}

func ComparePassword(hashedPwd string, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pass)); err != nil {
		return err
	}

	return nil
}
