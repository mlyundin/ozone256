package main

import (
	"log"
	"net/http"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/handlers/addtocart"
	"route256/checkout/internal/handlers/deletefromcart"
	"route256/checkout/internal/handlers/listcart"
	"route256/checkout/internal/handlers/purchase"
	"route256/libs/srvwrapper"
)

const port = ":8080"

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	lomsClient := loms.New(config.ConfigData.Services.Loms)
	productClient := products.New(config.ConfigData.Services.Products.Url, config.ConfigData.Services.Products.Token)
	businessLogic := domain.New(lomsClient, productClient)

	addToCartHandler := addtocart.New(businessLogic)
	deleteFromCartHandler := deletefromcart.New(businessLogic)
	listCartHandler := listcart.New(businessLogic)
	purchaseHandler := purchase.New(businessLogic)

	http.Handle("/addToCart", srvwrapper.New(addToCartHandler.Handle))
	http.Handle("/deleteFromCart", srvwrapper.New(deleteFromCartHandler.Handle))
	http.Handle("/listCart", srvwrapper.New(listCartHandler.Handle))
	http.Handle("/purchase", srvwrapper.New(purchaseHandler.Handle))

	log.Println("listening http at", port)
	err = http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}
