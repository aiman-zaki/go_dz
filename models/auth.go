package models

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

// ValidCredential :
// swagger:response validCredential
type ValidCredential struct {
	// in: body
	Body struct {
		//the success message
		Message string `json:"message"`
		// the credential given once successfully logined
		Auth *Auth `json:"auth"`
	}
}

// InvalidCredential :
// swagger:response invalidCredential
type InvalidCredential struct {
	// in: body
	Body struct {
		//the error message
		Message string `json:"message"`
	}
}

// Auth represents the Authentication Model for this application
//
// swagger:model
type Auth struct {
	// id for user
	// readOnly: true
	ID int64 `json:"id"`
	// email for auth
	// required: true
	Email string `json:"email"`
	// password for auth
	// required: true
	Password string `json:"password"`
	// accesstoken (jwt) expired (1 hours)
	// readOnly: true
	AcessToken string `pg:"-" json:"access_token"`
	// refreshtoken (jwt) expired (15 days)
	// readOnly: true
	RefreshToken string `pg:"-" json:"refresh_token"`
	// the role for this user
	// not required during login
	RoleID int64 `json:"role_id"`
	// swagger:ignore
	Role *Role `json:"role" pg:"fk:role_id"`
}

// GenerateToken : Generate JWT Token
func (auth Auth) GenerateToken(a *Auth) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user": auth.Email})
	a.AcessToken = tokenString
}

// HashAndSalt : generate hashed password
func (auth Auth) HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func (auth Auth) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
