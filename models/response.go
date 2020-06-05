package models

// NotAuthorized :
// swagger:response notAthorized
type NotAuthorized struct {
	// in: body
	Body struct {
		Message string
	}
}

// ErrorNotFound :
// swagger:response errorNotFound
type ErrorNotFound struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// DataAlreadyExisted :
// swagger:response dataAlreadyExisted
type DataAlreadyExisted struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}
