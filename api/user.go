package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github/beat-kuliah/sip_pad_backend/db/sqlc"
	"github/beat-kuliah/sip_pad_backend/utils"
	"net/http"
	"time"
)

type User struct {
	server *Server
}

func (u User) router(server *Server) {
	u.server = server

	serverGroup := server.router.Group("/users", AuthenticatedMiddleware())
	serverGroup.GET("", u.listUsers)
	serverGroup.GET("me", u.getLoggedInUser)
	serverGroup.PATCH("name", u.updateName)
}

func (u *User) listUsers(c *gin.Context) {
	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  10,
	}
	users, err := u.server.queries.ListUsers(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUsers := []UserResponse{}

	for _, v := range users {
		n := UserResponse{}.toUserResponse(&v)
		newUsers = append(newUsers, *n)
	}

	c.JSON(http.StatusOK, newUsers)
}

func (u *User) getLoggedInUser(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	user, err := u.server.queries.GetUserByID(context.Background(), userId)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized to access resources"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{}.toUserResponse(&user))
}

type UpdateNameType struct {
	Name string `json:"name" binding:"required"`
}

func (u *User) updateName(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	var userInfo UpdateNameType

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateNameParams{
		ID:        userId,
		Name:      userInfo.Name,
		UpdatedAt: time.Now(),
	}

	user, err := u.server.queries.UpdateName(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{}.toUserResponse(&user))
}

type RoleResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type UserResponse struct {
	ID        int64         `json:"id"`
	Username  string        `json:"username"`
	Name      string        `json:"name"`
	Role      *RoleResponse `json:"role,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (u UserResponse) toUserResponse(user *db.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Role:      nil,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u UserResponse) toUserResponseWithRole(row db.GetUserWithRoleRow) *UserResponse {
	var role *RoleResponse
	if row.RoleID.Valid {
		role = &RoleResponse{
			ID:          row.RoleID.Int64,
			Name:        row.RoleName.String,
			Description: row.RoleDescription.String,
		}
	}
	return &UserResponse{
		ID:        row.ID,
		Username:  row.Username,
		Name:      row.Name,
		Role:      role,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
