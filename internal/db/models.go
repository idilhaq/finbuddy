package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string
	Name      string
	Role      string `gorm:"default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Expense struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ExpenseType string    `json:"expense_type"`
	Category    string    `json:"category"`
	Amount      int       `json:"amount"`
	Note        string    `json:"note"`
	Date        string    `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonthlyPlan is a user's monthly budget plan
type MonthlyPlan struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	Month     time.Time  `gorm:"not null"` // e.g. 2025-06-01
	Items     []PlanItem `gorm:"foreignKey:MonthlyPlanID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PlanItem represents a budget item (e.g. groceries, gasoline)
type PlanItem struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MonthlyPlanID  uuid.UUID `gorm:"type:uuid;not null"`
	Title          string    `gorm:"not null"`
	Group          string    `gorm:"not null"` // Needs, Wants, Savings
	Category       string    `gorm:"not null"`
	Amount         int       `gorm:"not null"`
	Note           string
	FrequencyValue int       // e.g. 5
	FrequencyUnit  string    // e.g. times, days
	FrequencyRate  int       // e.g. 40000 (amount per unit)
	PocketID       *uuid.UUID
	GoalID         *uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Pocket groups multiple PlanItems together logically
type Pocket struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"not null"`
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// InvestmentGoal links PlanItems to long-term saving targets
type InvestmentGoal struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	Title      string    `gorm:"not null"`
	Target     int       `gorm:"not null"`
	TargetDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}