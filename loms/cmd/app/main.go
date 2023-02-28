package main

import (
	"log"
	"net/http"
	"route256/libs/srvwrapper"
	"route256/loms/internal/handlers"
)

const port = ":8081"

func main() {
	handler := handlers.New()
	http.Handle("/stocks", srvwrapper.New(handler.Stocks))
	http.Handle("/createOrder", srvwrapper.New(handler.CreateOrder))
	http.Handle("/listOrder", srvwrapper.New(handler.ListOrder))
	http.Handle("/cancelOrder", srvwrapper.New(handler.CancelOrder))
	http.Handle("/orderPayed", srvwrapper.New(handler.OrderPayed))

	log.Println("listening http at", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}
