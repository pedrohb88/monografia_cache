package main

import (
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	invoicesCache "monografia/cache/invoices"
	paymentsCache "monografia/cache/payments"
	"monografia/lib/cache"
	"monografia/lib/database"
	srv "monografia/service"
	"monografia/store/invoices"
	"monografia/store/payments"
	"monografia/transport"
	"monografia/transport/entity"
)

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("error loading config; %s", err.Error())
			return
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
	paymentsStore := paymentsCache.New(cache, payments.New(&db))
	invoicesStore := invoicesCache.New(cache, invoices.New(&db))

	// Services
	service := srv.New(paymentsStore, invoicesStore)

	entity := entity.New(service)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("\nfailed to listen: %v", err)
	}

	s := transport.NewServer(service, entity)

	log.Printf("\nserver listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("\nfailed to serve: %v", err)
	}
}
