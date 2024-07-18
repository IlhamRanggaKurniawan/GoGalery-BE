package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (*string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashedPasswordSTR := string(hashedPassword)

	return &hashedPasswordSTR, nil
}

func ComparePassword(hashedPassword string, password string) error{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}