package main

import (
	"context"
	"filtering_pagination/handler"
	"filtering_pagination/repositories"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	// Connect to PostgreSQL
	db, err := repositories.Connect(ctx, repositories.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1710519",
		Dbname:   "fashionstore",
	})
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	// Initialize layers
	repo := repositories.NewFashionStoreRepository(db)
	handler := handler.NewFashionHandler(repo)

	r := chi.NewRouter()

r.Route("/FashionStore/v1", func(r chi.Router) {
    r.Get("/filter", handler.Filter)
})


	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}