package db

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Category  string    `json:"category"`
	Amount    int       `json:"amount"`
	Note      string    `json:"note"`
	Date      string    `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MonthlyPlan struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID
	Month     string `gorm:"uniqueIndex"` // e.g. "2025-05"
	Needs     int
	Wants     int
	Savings   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Saving struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID
	Amount    int
	Note      string
	Date      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string
	Name      string
	Role      string `gorm:"default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
