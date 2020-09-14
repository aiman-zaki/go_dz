package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductCategory struct {
	ID          uuid.UUID `json:"id" dt:"id" pg:"type:uuid"`
	Category    string    `json:"category" `
	CreatedDate time.Time `json:"created_date" pg:"default:CURRENT_TIMESTAMP"`
	UpdatedDate time.Time `json:"date_updated" pg:"default:CURRENT_TIMESTAMP"`
}
