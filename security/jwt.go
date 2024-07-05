package security

// import (
// 	"time"
// 	"github.com/golang-jwt/jwt/v4"
// )

// var (
// 	JwtSecretKey     = []byte("ceit_")
// 	JwtPartnerSecret = []byte("jiv313a2")
// 	//jwtSigningMethod = jwt.SigningMethodHS256.Name
// )

// func NewAccessToken(userId string) (string, error) {
// 	claims := jwt.StandardClaims{
// 		Id:        userId,
// 		Issuer:    userId,
// 		IssuedAt:  time.Now().Unix(),
// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
// 	}
// 	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedString, err := withClaims.SignedString(JwtSecretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedString, nil
// }
