package models

import (
	"errors"
	"log"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg/v9"
	"golang.org/x/crypto/bcrypt"
)

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

type AuthWrapper struct {
	Auth Auth
	User User
}

func (aw *AuthWrapper) Register() error {
	db := pg.Connect(services.PgOptions())
	count, err := db.Model(&aw.Auth).
		Where("email = ?", aw.Auth.Email).
		Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("Account Existed")
	}
	hashed := Auth.HashAndSalt(Auth{}, []byte(aw.Auth.Password))
	aw.Auth.Password = hashed
	db.Insert(&aw.Auth)
	aw.User.ID = aw.Auth.ID
	db.Insert(&aw.User)
	return nil

}

func (aw *AuthWrapper) Login() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	plainPassword := aw.Auth.Password
	count, err := db.Model(&aw.Auth).
		Where("email = ?", aw.Auth.Email).
		SelectAndCount()
	if err != nil {
		return err
	}
	if count > 0 {
		valid := Auth.ComparePasswords(Auth{}, aw.Auth.Password, []byte(plainPassword))
		if valid {
			err := db.Model(&aw.User).
				Where(`"user"."id" = ?`, aw.Auth.ID).
				Relation("Auth").
				Select()
			if err != nil {
				return err
			}
			Auth.GenerateToken(Auth{}, &aw.Auth)
		}
	}
	return nil
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
