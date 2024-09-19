package handler

import (
	"gateway/proto/income"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new income transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   transaction  body      income.TransactionInfo  true  "Transaction info"
// @Success 201 {object} income.TransactionID
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [post]
func (h *Handler) CreateTransaction(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var createIncome income.TransactionInfo
	if err := protojson.Unmarshal(bytes, &createIncome); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}
	trID, err := h.Income.CreateTransaction(c.Request.Context(), &createIncome)
	if err != nil {
		log.Println("Error creating income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to create income"})
		return
	}

	c.IndentedJSON(201, trID)
}

// UpdateTransactionByID godoc
// @Summary Update a transaction by ID
// @Description Update an existing transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   transaction  body      income.TransactionWithID  true  "Transaction with ID"
// @Success 200 {object} income.TransactionID
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [put]
func (h *Handler) UpdateTransactionByID(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var TransactionWithID income.TransactionWithID
	if err := protojson.Unmarshal(bytes, &TransactionWithID); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}

	tres, err := h.Income.UpdateTransactionByID(c.Request.Context(), &TransactionWithID)
	if err != nil {
		log.Println("Error updating income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to update income"})
		return
	}

	c.IndentedJSON(200, tres)
}

// DeleteTransactionByID godoc
// @Summary Delete a transaction by ID
// @Description Delete a transaction by its ID
// @Tags transactions
// @Produce  json
// @Param   id  path   string  true  "Transaction ID"
// @Success 200 {object} income.TransactionID
// @Failure 500 {object} map[string]string
// @Router /transactions/{id} [delete]
func (h *Handler) DeleteTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var trId income.TransactionID
	trId.Id = id
	tres, err := h.Income.DeleteTransactionByID(c.Request.Context(), &trId)
	if err != nil {
		log.Println("Error deleted income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, tres)
}

// GetTransactionByID godoc
// @Summary Get a transaction by ID
// @Description Retrieve a transaction by its ID
// @Tags transactions
// @Produce  json
// @Param   id  path   string  true  "Transaction ID"
// @Success 200 {object} income.TransactionWithID
// @Failure 500 {object} map[string]string
// @Router /transactions/{id} [get]
func (h *Handler) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var trId income.TransactionID
	trId.Id = id
	transactionWithID, err := h.Income.GetTransactionByID(c.Request.Context(), &trId)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, transactionWithID)
}

// GetTransactionsByCategory godoc
// @Summary Get transactions by category
// @Description Retrieve transactions based on a specific category
// @Tags transactions
// @Produce  json
// @Param   category  path   string  true  "Transaction Category"
// @Success 200 {object} income.ListTransactions
// @Failure 500 {object} map[string]string
// @Router /transactions/category/{category} [get]
func (h *Handler) GetTransactionsByCategory(c *gin.Context) {
	category := c.Param("categroy")
	var trCategory income.TransactionCategory
	trCategory.Category = category
	list, err := h.Income.GetTransactionsByCategory(c.Request.Context(), &trCategory)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to get income"})
		return
	}

	c.IndentedJSON(200, list)
}

// GetTransactionByDate godoc
// @Summary Get transactions by date
// @Description Retrieve transactions based on a specific date
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   date  body      income.TransactionDate  true  "Transaction date"
// @Success 200 {object} income.ListTransactions
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/date [post]
func (h *Handler) GetTransactionByDate(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var trDate income.TransactionDate
	if err := protojson.Unmarshal(bytes, &trDate); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}
	lst, err := h.Income.GetTransactionByDate(c.Request.Context(), &trDate)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, lst)
}
