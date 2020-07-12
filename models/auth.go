package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Auth represents the Authentication Model for this application
//
// swagger:model
type Auth struct {
	// id for user
	// readOnly: true
	ID uuid.UUID `json:"id" pg:"type:uuid"`
	// email for auth
	// required: true
	Email string `json:"email"`
	// password for auth
	// required: true
	Username string `json:"username"`
	Password string `json:"password"`
	// accesstoken (jwt) expired (1 hours)
	// readOnly: true
	AcessToken string `pg:"-" json:"access_token"`
	// refreshtoken (jwt) expired (15 days)
	// readOnly: true
	RefreshToken string `pg:"-" json:"refresh_token"`
	// the role for this user
	// not required during login
	UserID      uuid.UUID `json:"user_id" pg:"type:uuid"`
	User        User      `json:"user" pg:"fk:user_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type AuthWrapper struct {
	Auth Auth
	User User
}

func (aw *AuthWrapper) Register() error {
	var err error
	db := pg.Connect(services.PgOptions())
	count, err := db.Model(&Auth{}).
		Where("email = ?", aw.Auth.Email).
		Where("username = ?", aw.Auth.Username).
		Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("Account Existed")
	}
	hashed := Auth.HashAndSalt(Auth{}, []byte(aw.Auth.Password))
	aw.Auth.Password = hashed
	aw.Auth.DateCreated = time.Now()
	aw.Auth.DateUpdated = time.Now()
	aw.Auth.ID = uuid.New()
	fmt.Println(aw)
	err = db.Insert(&aw.Auth)
	if err != nil {
		return err
	}
	return nil

}

func (aw *AuthWrapper) Login() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	plainPassword := aw.Auth.Password
	count, err := db.Model(&Auth{}).
		Where("email = ?", aw.Auth.Email).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		err := db.Model(&aw.Auth).Where("email = ?", aw.Auth.Email).Select()
		if err != nil {
			return err
		}
		valid := Auth.ComparePasswords(Auth{}, aw.Auth.Password, []byte(plainPassword))
		if !valid {
			return errors.New("Invalid Credential")
		}
		err = db.Model(&aw.Auth).
			Relation("User").
			Select()
		aw.GenerateToken()
		if err != nil {
			return err
		}
	} else {
		return errors.New("No Data")
	}
	return nil
}

// GenerateToken : Generate JWT Token
func (aw *AuthWrapper) GenerateToken() {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user": aw.Auth.Email})
	aw.Auth.AcessToken = tokenString
	fmt.Println(aw.Auth.AcessToken)
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
