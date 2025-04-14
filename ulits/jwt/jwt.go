package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
)


type Jwt struct {}
var jwtKey = []byte("1122390ifm")
func (j *Jwt) CreateToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,                    // Subject (user identifier)
		"iss": "todo-app",                  // Issuer
		"aud": "user",                   // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                 // Issued at
	})
tokenString, err := claims.SignedString(jwtKey)
if err != nil {
	return "", fmt.Errorf("failed to sign token: %w", err)
}
// Print information about the created token
fmt.Printf("Token claims added: %+v\n", claims)
return tokenString, nil
}

func (j *Jwt) VerifyToken(tokenString string) (string, error) {
    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validate the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })

    if err != nil {
        return "", fmt.Errorf("failed to parse token: %w", err)
    }

    // Validate the token claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        email := claims["sub"].(string) // Extract the email from the claims
        return email, nil
    }

    return "", fmt.Errorf("invalid token")
}