package main

import (
	"demo-hex-go/internal/adapter/db"
	"demo-hex-go/internal/adapter/db/repostory"
	"demo-hex-go/internal/adapter/handler"
	"demo-hex-go/internal/core/port"
	"demo-hex-go/internal/core/service"
	"flag"
	"fmt"
)

var (
	repo = flag.String("db", "postgres", "Database for storing messages")
	svc  port.ProductService
)

func main() {
	flag.Parse()

	fmt.Printf("Application running using %s\n", *repo)
	switch *repo {
	case "mongodb":
		setup := db.NewMongoRepository()
		repo := repostory.NewMongoRepository(setup)
		svc = service.NewProductMongoService(repo)
	default:
		setup := db.NewPostgresRepository()
		repo := repostory.NewPostgresRepository(setup)
		svc = service.NewProductService(repo)

	}

	fmt.Println("starting server on port 8081...")
	router := handler.RouteInit(svc)

	router.Listen(":8081")

}
