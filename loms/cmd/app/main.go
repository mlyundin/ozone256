package main

import (
	"log"
	"net/http"
	"route256/libs/srvwrapper"
	"route256/loms/internal/handlers"
)

const port = ":8081"

func main() {
	handler := handlers.NewHandler()
	http.Handle("/stocks", srvwrapper.New(handler.HandleStocks))
	http.Handle("/createOrder", srvwrapper.New(handler.HandleCreateOrder))
	http.Handle("/listOrder", srvwrapper.New(handler.HandleListOrder))
	http.Handle("/cancelOrder", srvwrapper.New(handler.HandleCancelOrder))
	http.Handle("/orderPayed", srvwrapper.New(handler.HandleOrderPayed))

	log.Println("listening http at", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}
