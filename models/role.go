package models

// SuccessRole :
// swagger:response successRole
type SuccessRole struct {
	// in: body
	Body struct {
		//the success message
		Message string `json:"message"`
		Role    *Role  `json:"role"`
	}
}

// SuccessRoleArray :
// swagger:response successRoleArray
type SuccessRoleArray struct {
	// in: body
	Body struct {
		//the success message
		Message string  `json:"message"`
		Role    []*Role `json:"roles"`
	}
}

// Role : model for the role existes in the system
// swagger:model
type Role struct {
	// id for role
	// readOnly: true
	ID int64 `json:"id"`
	// the role
	Role string `json:"role"`
}
