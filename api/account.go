package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github/beat-kuliah/finbest_backend/db/sqlc"
	"net/http"
)

type Account struct {
	server *Server
}

func (a Account) router(server *Server) {
	a.server = server

	serverGroup := server.router.Group("/account", AuthenticatedMiddleware())
	serverGroup.POST("create-account", a.createAccount)
}

type AccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (a Account) createAccount(c *gin.Context) {
	value, exists := c.Get("account")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to access resource"})
		return
	}

	userId, ok := value.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Encountered an issue"})
		return
	}

	acc := new(AccountRequest)

	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateAccountParams{
		UserID:   int32(userId),
		Currency: acc.Currency,
	}

	account, err := a.server.queries.CreateAccount(context.Background(), arg)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "You already have an account with this currency."})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}
