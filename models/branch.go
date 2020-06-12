package models

// BranchesResponse : List all branches
// swagger:response branches
type BranchesResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string    `json:"message"`
		Branch  *[]Branch `json:"branches"`
	}
}

// BranchResponse : List a branch
// swagger:response branch
type BranchResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string  `json:"message"`
		Branch  *Branch `json:"branch"`
	}
}

// Branch is a model bro
// swagger:model
type Branch struct {
	// readonly:true
	ID           int64      `json:"id"`
	CoordinateId int64      `json:"coordinate_id"`
	Name         string     `json:"name"`
	Address      string     `json:"address"`
	Coordinate   Coordinate `pg:"fk:coordinate_id" json:"coordinate"`
}
