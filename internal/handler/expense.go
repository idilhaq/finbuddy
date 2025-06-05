package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/idilhaq/finbuddy/internal/db"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateExpenseRequest struct {
	Amount   int    `json:"amount" binding:"required"`
	Category string `json:"category" binding:"required"`
	Date     string `json:"date" binding:"required"`
	Note     string `json:"note"`
}

// GetAllExpenses godoc
// @Summary      Get all expenses
// @Description  Returns a list of all expenses
// @Tags         Expenses
// @Produce      json
// @Success      200  {array}  db.Expense
// @Failure      500  {object}  ErrorResponse
// @Router       /api/expenses [get]
func GetAllExpenses(c *gin.Context) {
	var expenses []db.Expense
	query := db.DB

	// Optional filters
	date := c.Query("date")
	month := c.Query("month")
	category := c.Query("category")
	groupBy := c.Query("group_by")

	if date != "" {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err == nil {
			query = query.Where("date::date = ?", parsedDate)
		}
	}

	if month != "" {
		parsedMonth, err := time.Parse("2006-01", month)
		if err == nil {
			start := parsedMonth
			end := parsedMonth.AddDate(0, 1, 0)
			query = query.Where("date >= ? AND date < ?", start, end)
		}
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Grouping
	if groupBy == "date" {
		type DailySummary struct {
			Date  time.Time `json:"date"`
			Total int       `json:"total"`
		}

		rawSQL := `
		SELECT 
			date::date AS date,
			SUM(amount) AS total
		FROM expenses
		WHERE 1 = 1
	`

		var args []interface{}
		argCount := 1

		if date != "" {
			rawSQL += fmt.Sprintf(" AND date::date = $%d", argCount)
			parsedDate, _ := time.Parse("2006-01-02", date)
			args = append(args, parsedDate)
			argCount++
		}

		if month != "" {
			parsedMonth, _ := time.Parse("2006-01", month)
			start := parsedMonth
			end := parsedMonth.AddDate(0, 1, 0)
			rawSQL += fmt.Sprintf(" AND date >= $%d AND date < $%d", argCount, argCount+1)
			args = append(args, start, end)
			argCount += 2
		}

		if category != "" {
			rawSQL += fmt.Sprintf(" AND category = $%d", argCount)
			args = append(args, category)
			argCount++
		}

		rawSQL += " GROUP BY date::date ORDER BY date"

		var results []DailySummary
		if err := db.DB.Raw(rawSQL, args...).Scan(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, results)
		return
	}

	// Normal query
	if err := query.Order("date DESC").Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

// GetAllExpensesByUserID godoc
// @Summary      Get all expenses by user ID
// @Description  Returns a list of all expenses for a specific user
// @Tags         Expenses
// @Produce      json
// @Param        user_id  query     string  true  "User ID"
// @Success      200      {array}   db.Expense
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /api/expenses/user [get]
func GetAllExpensesByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var expenses []db.Expense
	query := db.DB.Where("user_id = ?", userID)

	// Optional filters
	date := c.Query("date")
	month := c.Query("month")
	category := c.Query("category")
	groupBy := c.Query("group_by")

	if date != "" {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err == nil {
			query = query.Where("date::date = ?", parsedDate)
		}
	}

	if month != "" {
		parsedMonth, err := time.Parse("2006-01", month)
		if err == nil {
			start := parsedMonth
			end := parsedMonth.AddDate(0, 1, 0)
			query = query.Where("date >= ? AND date < ?", start, end)
		}
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Grouping
	if groupBy == "date" {
		type DailySummary struct {
			Date  time.Time `json:"date"`
			Total int       `json:"total"`
		}

		rawSQL := `
		SELECT 
			date::date AS date,
			SUM(amount) AS total
		FROM expenses
		WHERE user_id = ?
	`
		var args []interface{}
		args = append(args, userID)
		argCount := 2

		if date != "" {
			rawSQL += fmt.Sprintf(" AND date::date = $%d", argCount)
			parsedDate, _ := time.Parse("2006-01-02", date)
			args = append(args, parsedDate)
			argCount++
		}

		if month != "" {
			parsedMonth, _ := time.Parse("2006-01", month)
			start := parsedMonth
			end := parsedMonth.AddDate(0, 1, 0)
			rawSQL += fmt.Sprintf(" AND date >= $%d AND date < $%d", argCount, argCount+1)
			args = append(args, start, end)
			argCount += 2
		}

		if category != "" {
			rawSQL += fmt.Sprintf(" AND category = $%d", argCount)
			args = append(args, category)
			argCount++
		}

		rawSQL += " GROUP BY date::date ORDER BY date"

		var results []DailySummary
		if err := db.DB.Raw(rawSQL, args...).Scan(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, results)
		return
	}

	// Normal query
	if err := query.Order("date DESC").Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

// CreateExpense godoc
// @Summary      Create a new expense
// @Description  Add a new expense record
// @Tags         Expenses
// @Accept       json
// @Produce      json
// @Param        expense  body      db.Expense  true  "Expense data"
// @Success      201      {object}  db.Expense
// @Failure      400      {object}  ErrorResponse
// @Failure 	 500      {object}  ErrorResponse
// @Router       /api/expenses [post]
func CreateExpense(c *gin.Context) {
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var createExpense CreateExpenseRequest
	if err := c.ShouldBindJSON(&createExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDStr, ok := userIDRaw.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid UUID"})
		return
	}

	expense := db.Expense{
		ID:        uuid.New(),
		UserID:    uid,
		Category:  createExpense.Category,
		Amount:    createExpense.Amount,
		Note:      createExpense.Note,
		Date:      createExpense.Date,
		CreatedAt: time.Now(),
	}

	if err := db.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, expense)
}

// GetExpenseByID godoc
// @Summary      Get a single expense
// @Description  Retrieve expense by ID
// @Tags         Expenses
// @Produce      json
// @Param        id   path      string  true  "Expense ID"
// @Success      200  {object}  db.Expense
// @Failure      404  {object}  ErrorResponse
// @Router       /api/expenses/{id} [get]
func GetExpenseByID(c *gin.Context) {
	id := c.Param("id")
	var expense db.Expense
	if err := db.DB.First(&expense, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}
	c.JSON(http.StatusOK, expense)
}

// UpdateExpense godoc
// @Summary      Update an expense
// @Description  Update a specific expense by ID
// @Tags         Expenses
// @Accept       json
// @Produce      json
// @Param        id       path      string      true  "Expense ID"
// @Param        expense  body      db.Expense  true  "Updated expense"
// @Success      200      {object}  db.Expense
// @Failure      400      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /api/expenses/{id} [put]
func UpdateExpense(c *gin.Context) {
	id := c.Param("id")
	var expense db.Expense
	if err := db.DB.First(&expense, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Save(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expense)
}

// DeleteExpense godoc
// @Summary      Delete an expense
// @Description  Remove an expense by ID
// @Tags         Expenses
// @Produce      json
// @Param        id   path      string  true  "Expense ID"
// @Success      204  {string}  string  "No Content"
// @Failure      500  {object}  ErrorResponse
// @Router       /api/expenses/{id} [delete]
func DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&db.Expense{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
