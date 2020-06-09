package models

type Stock struct {
	ID        int64    `json:"id"`
	BranchId  int64    `json:"branch_id"`
	ProductId int64    `json:"product_id"`
	UnitId    int64    `json:"unit_id"`
	Amount    int64    `json:"amount"`
	Unit      *Unit    `pg:"fk:unit_id"`
	Branch    *Branch  `pg:"fk:branch_id"`
	Product   *Product `pg:"fk:product_id"`
}
