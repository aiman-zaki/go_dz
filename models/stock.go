package models

type Stock struct {
	ID           int64
	UnitBundleId int64
	StoreId      int64
	ProductId    int64
	UnitBundle   *UnitBundle `pg:"fk:unit_bundle_id"`
	Store        *Store      `pg:"fk:store_id"`
	Product      *Product    `pg:"fk:product_id"`
}
