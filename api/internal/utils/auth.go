package utils

import (
	"crypto/ed25519"
	"errors"
	"time"

	"plant-manager/internal/types"

	"github.com/pascaldekloe/jwt"
)

var (
	prv ed25519.PrivateKey
	pub ed25519.PublicKey
)

func init() {
	var err error
	pub, prv, err = ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
}

func GetNewClaims(subject string, set map[string]interface{}) *jwt.Claims {
	return &jwt.Claims{
		Registered: jwt.Registered{
			Subject: subject,
		},
		Set: set,
	}
}

func GetSignedToken(claims *jwt.Claims) (string, error) {
	now := time.Now().Round(time.Second)

	claims.Registered.Issued = jwt.NewNumericTime(now)
	claims.Registered.Expires = jwt.NewNumericTime(now.Add(7 * time.Hour * 24))
	claims.NotBefore = jwt.NewNumericTime(now.Add(-time.Second))

	token, err := claims.EdDSASign(prv)
	if err != nil {
		return "", nil
	}

	return string(token), nil
}

func IsJWTValid(token string) bool {
	if _, err := jwt.EdDSACheck([]byte(token), pub); err != nil {
		return false
	}
	return true
}

func GetDataFromToken(token string) (*types.User, error) {
	claims, err := jwt.EdDSACheck([]byte(token), pub)
	if err != nil {
		return nil, err
	}

	userDataErr := errors.New("token is valid but user data missing or corrupt")
	userData, ok := claims.Set["user"].(map[string]interface{})
	if !ok {
		return nil, userDataErr
	}

	user := new(types.User)

	id, ok := userData["id"].(float64)
	if !ok {
		return nil, userDataErr
	}
	user.ID = int64(id)

	email, ok := userData["email"].(string)
	if !ok {
		return nil, userDataErr
	}
	user.Email = email

	firstName, ok := userData["first_name"].(string)
	if !ok {
		return nil, userDataErr
	}
	user.FirstName = firstName

	lastName, ok := userData["last_name"].(string)
	if !ok {
		return nil, userDataErr
	}
	user.LastName = lastName

	customerID, ok := userData["customer_id"].(float64)
	if !ok {
		return nil, userDataErr
	}
	user.CustomerID = int64(customerID)

	return user, nil
}
