package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Andres-Salamanca/microcourse/handlers"
	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)


	hp := handlers.NewProducts(logger)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/",hp.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", hp.UpdateProducts)
	putRouter.Use(hp.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/",hp.AddProduct)
	postRouter.Use(hp.MiddlewareProductValidation)

	ser := &http.Server{Addr: ":9090",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}


	go func() {
		logger.Println("Starting server on port 9090")
		err:=ser.ListenAndServe()
		if err!= nil {
			logger.Fatal(err)
			os.Exit(1)
		}
	}()
	
	sigchan := make(chan os.Signal,1)
	signal.Notify(sigchan,os.Interrupt)
	signal.Notify(sigchan,os.Kill)

	sig :=  <- sigchan
	logger.Println("Recived terminated gracefully shutdown",sig)
	tc,_ := context.WithTimeout(context.Background(),30* time.Second)
	ser.Shutdown(tc)

}
