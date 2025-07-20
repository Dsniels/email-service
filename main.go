package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/dsniels/email-service/internal/api"
	"github.com/dsniels/email-service/internal/api/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := flag.String("port", ":3000", "the port")
	flag.Parse()

	app, err := api.NewApp()
	if err != nil {
		panic(err)
	}

	go app.Rab.StartConsuming(context.Background(), "Mail")

	api := router.InitRoutes(app)
	server := &http.Server{
		Handler: router.RecoverMiddleware(api),
		Addr:    *port,
	}
	log.Println("Server running on port: ", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
