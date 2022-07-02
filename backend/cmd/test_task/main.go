package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"test/internal/api"
	"test/internal/pkg"
	"test/internal/pkg/database"
)

func main() {
	port := "8080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	db, err := sql.Open("postgres", "postgresql://postgres:123@localhost?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dao := database.CreateNewDao(db)
	fs := pkg.CreateNewFileService(dao)
	ks := pkg.CreateNewKeyService(dao)

	c := api.CreateNewController(fs, ks)
	cf := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Access-Control-Allow-Origin"},
	})
	log.Fatal(http.ListenAndServe(":"+port, cf.Handler(c.Router)))
}
