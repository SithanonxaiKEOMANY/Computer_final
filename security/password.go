package security

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, Password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(Password))
}

//func HashPassword(password string) (string, error) {
//	cost := bcrypt.DefaultCost
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
//	if err != nil {
//		return "", err
//	}
//	return string(hashedPassword), err
//}
//
//func CheckPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}
