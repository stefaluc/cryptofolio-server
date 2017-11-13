package main

import (
	// Launch server
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strconv"
	"time"

	"fmt"

	"github.com/stefaluc/cryptofolio-server/database"
	"github.com/stefaluc/cryptofolio-server/server"
)

const (
	defaultPort = 8080
	dbUser      = "cryptofolio"
	dbPassword  = "cryptofolio"
	dbName      = "cryptofolio"
)

func main() {
	// init DB
	var errdb error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	database.DBConn, errdb = sql.Open("postgres", dbinfo)
	if errdb != nil {
		panic(errdb)
	}
	defer database.DBConn.Close()

	// Environment vars.
	// Port.
	port := uint16(defaultPort)
	portStr := os.Getenv("PORT")
	if portStr != "" {
		p, err := strconv.ParseUint(portStr, 10, 16)
		if err != nil {
			fmt.Printf("invalid port: %s", err)
		}
		port = uint16(p)
	}

	// Server.
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
		Handler:        server.NewRouter(),
	}
	panic(srv.ListenAndServe())
}
