package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github/beat-kuliah/sip_pad_backend/db/sqlc"
	"github/beat-kuliah/sip_pad_backend/utils"
	"net/http"
)

type Auth struct {
	server *Server
}

func (a Auth) router(server *Server) {
	a.server = server

	serverGroup := server.router.Group("/auth")
	serverGroup.POST("login", a.login)
	serverGroup.POST("register", a.register)
}

type UserParams struct {
	Username string `json:"username" binding:"required"`
	RoleID   *int64 `json:"roleID"`
	Password string `json:"password" binding:"required"`
}

func (a *Auth) register(c *gin.Context) {
	var user UserParams
	eViewer := gValid.Validator(UserParams{})
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.HandlerError(err, c, eViewer)})
		return
	}

	hashedPassword, err := utils.GenerateHashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// set sql NullINT64
	var roleID sql.NullInt64
	if user.RoleID != nil {
		roleID = sql.NullInt64{Int64: *user.RoleID, Valid: true}
	} else {
		roleID = sql.NullInt64{Valid: false}
	}

	args := db.CreateUserParams{
		Username:       user.Username,
		Name:           user.Username,
		RoleID:         roleID,
		HashedPassword: hashedPassword,
	}

	newUser, err := a.server.queries.CreateUser(context.Background(), args)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newUserWithRole, err := a.server.queries.GetUserWithRole(context.Background(), newUser.ID)
	c.JSON(http.StatusCreated, UserResponse{}.toUserResponseWithRole(newUserWithRole))

}

func (a Auth) login(c *gin.Context) {
	user := new(UserParams)
	eViewer := gValid.Validator(UserParams{})

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.HandlerError(err, c, eViewer)})
		return
	}

	dbUser, err := a.server.queries.GetUserByUsername(context.Background(), user.Username)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect username or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := utils.VerifyPassword(user.Password, dbUser.HashedPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect username or password"})
		return
	}

	token, err := tokenController.CreateToken(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
