package main

import (
	"log"
	"net"
	"net/http"
	"server/internal/user"
	"server/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Starting server")
	log.Println("Server started")
	router := httprouter.New()
	log.Println("register user handler")

	start(router)
	handler := user.NewHandler(logger)
	handler.Register(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	listener, err := net.Listen("tcp", ":8080")
	logger.Info("Listening on", listener.Addr().String())
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := server.Serve(listener); err != nil {
		logger.Fatalln("Error serving: ", err)
	}
}
