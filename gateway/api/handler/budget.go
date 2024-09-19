package handler

import (
	"gateway/proto/budget"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateBudget godoc
// @Summary Create a new budget
// @Description Create a new budget with the provided information
// @Tags budget
// @Accept json
// @Produce json
// @Param budget body budget.BudgetInfo true "Create Budget"
// @Success 201 {object} budget.BudgetID
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Unable to create budget"
// @Router /budgets [post]
func (h *Handler) CreateBudget(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var createBudget budget.BudgetInfo
	if err := protojson.Unmarshal(bytes, &createBudget); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetID, err := h.Budget.CreateBudget(c.Request.Context(), &createBudget)
	if err != nil {
		log.Println("Error creating budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to create budget"})
		return
	}

	c.IndentedJSON(201, budgetID)
}

// UpdateBudgetAmount godoc
// @Summary Update the budget amount
// @Description Update the budget amount for a specific budget
// @Tags budget
// @Accept json
// @Produce json
// @Param budget body budget.BudgetUpdate true "Update Budget Amount"
// @Success 200 {object} budget.BudgetResponse
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Unable to update budget"
// @Router /budgets/update [put]
func (h *Handler) UpdateBudgetAmount(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var updateBudget budget.BudgetUpdate
	if err := protojson.Unmarshal(bytes, &updateBudget); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetResponse, err := h.Budget.UpdateBudgetAmount(c.Request.Context(), &updateBudget)
	if err != nil {
		log.Println("Error updating budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to update budget"})
		return
	}

	c.IndentedJSON(200, budgetResponse)
}

// GetBudgets godoc
// @Summary Get all budgets
// @Description Retrieve a list of all budgets
// @Tags budget
// @Accept json
// @Produce json
// @Success 200 {object} budget.BudgetWithID
// @Failure 500 {object} map[string]string "Unable to retrieve budgets"
// @Router /budgets [get]
func (h *Handler) GetBudgets(c *gin.Context) {
	var budgetEmpty budget.Empty
	budgetWithID, err := h.Budget.GetBudgets(c.Request.Context(), &budgetEmpty)
	if err != nil {
		log.Println("Error getting budgets:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to retrieve budgets"})
		return
	}

	c.IndentedJSON(200, budgetWithID)
}

// DeleteBudgetByID godoc
// @Summary Delete a budget by ID
// @Description Delete a budget by its ID
// @Tags budget
// @Accept json
// @Produce json
// @Param budget body budget.BudgetID true "Budget ID"
// @Success 200 {object} budget.BudgetResponse
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Unable to delete budget"
// @Router /budgets/delete [delete]
func (h *Handler) DeleteBudgetByID(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var budgetID budget.BudgetID
	if err := protojson.Unmarshal(bytes, &budgetID); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetResponse, err := h.Budget.DeleteBudgetByID(c.Request.Context(), &budgetID)
	if err != nil {
		log.Println("Error deleting budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete budget"})
		return
	}

	c.IndentedJSON(200, budgetResponse)
}
