package api

import (
	"context"
	"github.com/gin-gonic/gin"
	db "github/beat-kuliah/finbest_backend/db/sqlc"
	"net/http"
)

type User struct {
	server *Server
}

func (u User) router(server *Server) {
	u.server = server

	serverGroup := server.router.Group("/users")
	serverGroup.GET("", u.listUsers)
	serverGroup.POST("", u.createUser)
}

type UserParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) createUser(c *gin.Context) {
	var user UserParams

	c.ShouldBindJSON(&user)
}

func (u *User) listUsers(c *gin.Context) {
	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  10,
	}
	users, err := u.server.queries.ListUsers(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, users)
}
