package db

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Category  string    `json:"category"`
	Amount    int       `json:"amount"`
	Note      string    `json:"note"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
