package main

import (
	// Launch server
	"net/http"
	"os"
	"strconv"
	"time"

	"fmt"

	"github.com/stefaluc/cryptofolio-server/server"
)

const (
	defaultPort = 8080
)

func main() {
	//dbprovider, err := database.InitDB()
	//if err != nil {
	//panic(err)
	//}
	//defer dbprovider.Close()

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
