package models

// StockResponse :
// swagger:response stock
type StockResponse struct {
	//in:body
	Body struct {
		Message string `json:"message"`
		Stock   *Stock `json:"stock"`
	}
}

// StocksResponse :
// swagger:response stocks
type StocksResponse struct {
	//in:body
	Body struct {
		Message string   `json:"message"`
		Stock   []*Stock `json:"stocks"`
	}
}

// Stock : model
// swagger:model
type Stock struct {
	// readOnly:true
	ID        int64 `json:"id"`
	BranchID  int64 `json:"branch_id"`
	ProductID int64 `json:"product_id"`
	UnitID    int64 `json:"unit_id"`
	Amount    int64 `json:"amount"`
	// readOnly:true
	Unit Unit `pg:"fk:unit_id" json:"unit"`
	// readOnly:true
	Branch Branch `pg:"fk:branch_id" json:"branch"`
	// readOnly:true
	Product Product `pg:"fk:product_id" json:"product"`
}
