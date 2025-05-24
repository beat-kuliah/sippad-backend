package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golodash/galidator"
	_ "github.com/lib/pq"
	db "github/beat-kuliah/finbest_backend/db/sqlc"
	"github/beat-kuliah/finbest_backend/utils"
	"net/http"
)

type Server struct {
	queries *db.Store
	router  *gin.Engine
	config  *utils.Config
}

var tokenController *utils.JWTToken
var gValid = galidator.New().CustomMessages(
	galidator.Messages{
		"required": "This field is required.",
	},
)

func myCorsHandler() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	return cors.New(config)
}

func NewServer(envPath string) *Server {
	config, err := utils.LoadConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Could not load env config: %v", err))
	}

	conn, err := sql.Open(config.DBdriver, config.DB_source+config.DB_name+"?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	tokenController = utils.NewJWTToken(config)

	q := db.NewStore(conn)

	g := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", currencyValidator)
	}

	g.Use(myCorsHandler())

	return &Server{
		queries: q,
		router:  g,
		config:  config,
	}
}

func (s *Server) Start(port int) {
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Finbest!",
		})
	})

	User{}.router(s)
	Auth{}.router(s)

	s.router.Run(fmt.Sprintf(":%d", port))
}
