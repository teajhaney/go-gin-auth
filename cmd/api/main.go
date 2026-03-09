package main

import (
	"context"
	"go-auth/internal/app"
	httpserver "go-auth/internal/server"
	"log"
	"net/http"
	"time"
)



func main() {

	ctx:= context.Background()

	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}


	defer func (){
		if err := app.Close(ctx); err != nil {
			log.Fatalf("Failed to close app: %v", err)
		}
	}()
 


	router := httpserver.NewRouter(app)

	
	//standard go http server setup
	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
		ReadTimeout: 5*time.Second,
	}

	log.Printf("Starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil{
		if err == http.ErrServerClosed {
			log.Printf("Server closed")
			return
		}
		log.Fatalf("Server error: %s", err)
	} else {
		log.Printf("Server started")
	}

	}

	