// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/users/forgot-password": {
            "post": {
                "description": "This method is responsible for sending a password reset code to the user's email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USERS"
                ],
                "summary": "Request password reset",
                "parameters": [
                    {
                        "description": "User ID for reset",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.UserID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Code sent to your email",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Error retrieving email",
                        "schema": {}
                    },
                    "403": {
                        "description": "Permission Denied",
                        "schema": {}
                    },
                    "500": {
                        "description": "Error storing reset code or sending email",
                        "schema": {}
                    }
                }
            }
        },
        "/api/users/login": {
            "post": {
                "description": "This method is responsible for user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USERS"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful with token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Error while unmarshaling",
                        "schema": {}
                    },
                    "401": {
                        "description": "Invalid username or password",
                        "schema": {}
                    },
                    "403": {
                        "description": "Permission Denied",
                        "schema": {}
                    }
                }
            }
        },
        "/api/users/logout": {
            "post": {
                "description": "This method is responsible for logging out the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USERS"
                ],
                "summary": "User logout",
                "parameters": [
                    {
                        "description": "User ID for logout",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.UserID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Logout successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Error while Logginf out user ",
                        "schema": {}
                    },
                    "403": {
                        "description": "Permission Denied",
                        "schema": {}
                    }
                }
            }
        },
        "/api/users/register": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This method is responsible for registering new users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USERS"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body or error unmarshaling",
                        "schema": {}
                    },
                    "403": {
                        "description": "Permission Denied",
                        "schema": {}
                    },
                    "500": {
                        "description": "Error storing user data or sending confirmation email",
                        "schema": {}
                    }
                }
            }
        },
        "/api/users/update-password": {
            "post": {
                "description": "This method is responsible for updating the user's password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USERS"
                ],
                "summary": "Update user password",
                "parameters": [
                    {
                        "description": "New password details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.NewPass"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password updated successfully",
                        "schema": {
                            "$ref": "#/definitions/users.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Invalid input or user verification error",
                        "schema": {}
                    },
                    "403": {
                        "description": "Permission Denied",
                        "schema": {}
                    },
                    "500": {
                        "description": "Error while updating password",
                        "schema": {}
                    }
                }
            }
        },
        "/api/users/verification": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This method is responsible for verifying the user registration code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USERS"
                ],
                "summary": "Verify user registration code",
                "parameters": [
                    {
                        "description": "User code details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCode"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/users.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Error verifying code",
                        "schema": {}
                    },
                    "500": {
                        "description": "Error creating user or sending welcome email",
                        "schema": {}
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "description": "This method retrieves a user by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USERS"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User data retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/users.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Error while retrieving user",
                        "schema": {}
                    },
                    "403": {
                        "description": "Permission Denied",
                        "schema": {}
                    }
                }
            }
        },
        "/budgets": {
            "get": {
                "description": "Retrieve a list of all budgets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budget"
                ],
                "summary": "Get all budgets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/budget.BudgetWithID"
                        }
                    },
                    "500": {
                        "description": "Unable to retrieve budgets",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new budget with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budget"
                ],
                "summary": "Create a new budget",
                "parameters": [
                    {
                        "description": "Create Budget",
                        "name": "budget",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/budget.BudgetInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/budget.BudgetID"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Unable to create budget",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/budgets/delete": {
            "delete": {
                "description": "Delete a budget by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budget"
                ],
                "summary": "Delete a budget by ID",
                "parameters": [
                    {
                        "description": "Budget ID",
                        "name": "budget",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/budget.BudgetID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/budget.BudgetResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Unable to delete budget",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/budgets/update": {
            "put": {
                "description": "Update the budget amount for a specific budget",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budget"
                ],
                "summary": "Update the budget amount",
                "parameters": [
                    {
                        "description": "Update Budget Amount",
                        "name": "budget",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/budget.BudgetUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/budget.BudgetResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Unable to update budget",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/notifications": {
            "get": {
                "description": "Retrieve a list of all notifications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notifications"
                ],
                "summary": "Get all notifications",
                "responses": {
                    "200": {
                        "description": "List of notifications",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/notification.NotifyList"
                            }
                        }
                    },
                    "500": {
                        "description": "Unable to get notification",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/notifications/unread": {
            "get": {
                "description": "Retrieve a list of unread notifications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notifications"
                ],
                "summary": "Get unread notifications",
                "responses": {
                    "200": {
                        "description": "List of unread notifications",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/notification.NotifyList"
                            }
                        }
                    },
                    "500": {
                        "description": "Unable to get notification",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reports": {
            "get": {
                "description": "Retrieve the report based on the provided criteria",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "Get report",
                "responses": {
                    "200": {
                        "description": "Report details",
                        "schema": {
                            "$ref": "#/definitions/report.ReportResponse"
                        }
                    },
                    "500": {
                        "description": "Unable to get report",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "put": {
                "description": "Update an existing transaction by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Update a transaction by ID",
                "parameters": [
                    {
                        "description": "Transaction with ID",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/income.TransactionWithID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/income.TransactionID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new income transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Create a new transaction",
                "parameters": [
                    {
                        "description": "Transaction info",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/income.TransactionInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/income.TransactionID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/transactions/category/{category}": {
            "get": {
                "description": "Retrieve transactions based on a specific category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get transactions by category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction Category",
                        "name": "category",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/income.ListTransactions"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/transactions/date": {
            "post": {
                "description": "Retrieve transactions based on a specific date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get transactions by date",
                "parameters": [
                    {
                        "description": "Transaction date",
                        "name": "date",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/income.TransactionDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/income.ListTransactions"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/transactions/{id}": {
            "get": {
                "description": "Retrieve a transaction by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get a transaction by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/income.TransactionWithID"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a transaction by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Delete a transaction by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/income.TransactionID"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "budget.BudgetID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "budget.BudgetInfo": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "spent": {
                    "type": "number"
                }
            }
        },
        "budget.BudgetResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "budget.BudgetUpdate": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "budget.BudgetWithID": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "spent": {
                    "type": "number"
                }
            }
        },
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "income.ListTransactions": {
            "type": "object",
            "properties": {
                "listTransactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/income.TransactionWithID"
                    }
                }
            }
        },
        "income.TransactionDate": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "string"
                },
                "start": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "income.TransactionID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "income.TransactionInfo": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "income.TransactionWithID": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.NewPass": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserCode": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "notification.Notify": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "report": {
                    "$ref": "#/definitions/notification.Report"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "notification.NotifyList": {
            "type": "object",
            "properties": {
                "notifyList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/notification.Notify"
                    }
                }
            }
        },
        "notification.Report": {
            "type": "object",
            "properties": {
                "expenses": {
                    "type": "number"
                },
                "income": {
                    "type": "number"
                },
                "netSavings": {
                    "type": "number"
                }
            }
        },
        "report.Budget": {
            "type": "object",
            "properties": {
                "remainingBudget": {
                    "type": "number"
                },
                "totalAmount": {
                    "type": "number"
                },
                "totalSpent": {
                    "type": "number"
                }
            }
        },
        "report.ReportResponse": {
            "type": "object",
            "properties": {
                "budget": {
                    "$ref": "#/definitions/report.Budget"
                },
                "expenses": {
                    "type": "number"
                },
                "income": {
                    "type": "number"
                },
                "netSavings": {
                    "type": "number"
                }
            }
        },
        "users.UserID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "users.UserInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "users.UserLogin": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "PERSONAL FINANCE MANAGEMENT",
	Description:      "This swagger UI was created to manage personal finance",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
