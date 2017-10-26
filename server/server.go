package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/stefaluc/cryptofolio-server/handler"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowMethods("DELETE")
	r.Use(cors.New(config))

	r.
		POST("/auth", tonic.Handler(handler.Login, http.StatusOK)).
		POST("/signup", tonic.Handler(handler.SignUp, http.StatusOK))

	authenticated := r.Group("/auth")
	authenticated.
		POST("/balance", tonic.Handler(handler.InsertBalance, http.StatusOK)).
		POST("/balance/:balance", tonic.Handler(handler.InsertTransaction, http.StatusOK)).
		GET("/balances", tonic.Handler(handler.GetBalances, http.StatusOK)).
		GET("/transactions", tonic.Handler(handler.GetTransactions, http.StatusOK)).
		GET("/currencies", tonic.Handler(handler.GetCurrencies, http.StatusOK))

	return r
}
