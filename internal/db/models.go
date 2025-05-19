package db

import "time"

type Expense struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Category  string
	Amount    int
	Note      string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
