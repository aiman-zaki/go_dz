package models

import "time"

// UsersResponse : List all users
// swagger:response users
type UsersResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string  `json:"message"`
		User    []*User `json:"users"`
	}
}

// UserResponse : List all users
// swagger:response user
type UserResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string `json:"message"`
		User    *User  `json:"user"`
	}
}

// User represents the product for this application
//
// swagger:model
type User struct {
	// the id for this user
	// readOnly: true
	ID int64 `pg:"alias:auth_id" json:"id"`
	// swagger:ignore
	Auth *Auth `pg:"fk:auth_id"`
	// the first name for this user
	// required: true
	// min length: 3
	FirstName string `json:"first_name"`
	// the last name for this user
	// required: true
	// min length: 3
	LastName string `json:"last_name"`
	// the salary for this user
	Salary float64 `json:"salary"`
	// the dateCreated for this user
	DateCreated time.Time `json:"date_created" pg:"default:now()"`
	// the dateUpdated for this user
	DateUpdated time.Time `json:"date_updated" pg:"default:now()"`
}

func (u User) BeforeInsert() error {
	if u.DateCreated.IsZero() {
		u.DateCreated = time.Now()
	}

	return nil
}
