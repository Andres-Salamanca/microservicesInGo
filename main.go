package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Andres-Salamanca/microcourse/handlers"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)


	/*hw := handlers.NewHello(logger)
	hg := handlers.NewGoodBye(logger)*/
	hp := handlers.NewProducts(logger)
	sm := http.NewServeMux()
	/*sm.Handle("/", hw)
	sm.Handle("/goodbye", hg)*/

	sm.Handle("/", hp)

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
