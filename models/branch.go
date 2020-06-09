package models

type Branch struct {
	ID           int64      `json:"id"`
	CoordinateId int64      `json:"coordinate_id"`
	Name         string     `json:"name"`
	Address      string     `json:"address"`
	Coordinate   Coordinate `pg:"fk:coordinate_id" json:"coordinate"`
}
