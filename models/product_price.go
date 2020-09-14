package models

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	CreatedDate time.Time
	value       interface{}
}

type ProductPrice struct {
	ID         int64     `json:"id"`
	ProductID  uuid.UUID `json:"product_id" pg:"type:uuid"`
	Product    *Product  `json:"product" pg:"fk:product_id"`
	BranchID   uuid.UUID `json:"branch_id" pg:"type:uuid"`
	Branch     *Branch   `json:"branch" pg:"fk:branch_id"`
	CostPriceL float32   `json:"cost_priceL"`
	SellPriceL float32   `json:"sell_priceL"`
	CostPriceP struct{}  `json:"cost_priceP" pg:",array"`
	SellPriceP struct{}  `json:"sale_priceP" pg:",array"`
}

type ProductPriceWrapper struct {
	Single ProductPrice
	Array  []ProductPrice
}
