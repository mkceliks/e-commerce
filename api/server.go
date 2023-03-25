package api

import (
	db "github.com/e-commerce/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {

	server := &Server{
		store: store,
	}

	router := gin.Default()

	// router.POST("/accounts", server.createAccount)
	server.router = router

	return server
}
