package models

type RecordSupply struct {
	Id       int64
	SupplyId int64
	Supply   *Supply `pg:"fk:supply_id"`
}
