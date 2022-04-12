package main

import (
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	itemsCache "monografia/cache/items"
	ordersCache "monografia/cache/orders"
	productsCache "monografia/cache/products"
	"monografia/lib/cache"
	"monografia/lib/database"
	srv "monografia/service"
	"monografia/store/items"
	"monografia/store/orders"
	"monografia/store/payments"
	"monografia/store/products"
	"monografia/transport"
	"monografia/transport/entity"
)

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("error loading config; %s", err.Error())
		}
	}

	// Cache
	cache, err := cache.New()
	if err != nil {
		log.Fatal(err)
	}

	// Database
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	// Stores + cache
	ordersStore := ordersCache.New(cache, orders.New(&db))
	productsStore := productsCache.New(cache, products.New(&db))
	itemsStore := itemsCache.New(cache, items.New(&db))

	// Stores
	paymentsStore := payments.New()

	// Services
	service := srv.New(ordersStore, productsStore, itemsStore, paymentsStore)

	entity := entity.New(service)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("\nfailed to listen: %v", err)
	}

	s := transport.NewServer(service, entity)

	log.Printf("\nserver listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("\nfailed to serve: %v", err)
	}
}
