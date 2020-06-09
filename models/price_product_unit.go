package models

type PriceProductUnit struct {
	Id        string
	ProductId int64
	Product   *Product `pg:"fk:product_id"`
	UnitId    int64
	Unit      *Unit `pg:"fk:unit_id"`
	Price     float64
}
