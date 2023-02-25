package main

import (
	"log"
	"net/http"
	"route256/libs/srvwrapper"
	"route256/loms/internal/handlers"
)

const port = ":8081"

func main() {
	stocksHandler := handlers.NewHandler[handlers.StocksHandler]()
	createOrderHandler := handlers.NewHandler[handlers.CreateOrderHandler]()
	cancelOrderHandler := handlers.NewHandler[handlers.CancelOrderHandler]()
	listOrderHandler := handlers.NewHandler[handlers.ListOrderHandler]()
	oderPayedHandler := handlers.NewHandler[handlers.OrderPayedHandler]()

	http.Handle("/stocks", srvwrapper.New(stocksHandler.Handle))
	http.Handle("/createOrder", srvwrapper.New(createOrderHandler.Handle))
	http.Handle("/listOrder", srvwrapper.New(listOrderHandler.Handle))
	http.Handle("/cancelOrder", srvwrapper.New(cancelOrderHandler.Handle))
	http.Handle("/orderPayed", srvwrapper.New(oderPayedHandler.Handle))

	log.Println("listening http at", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}
