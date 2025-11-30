package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CryptoGu1/books-grpc-log/internal/config"
	"github.com/CryptoGu1/books-grpc-log/internal/repo"
	"github.com/CryptoGu1/books-grpc-log/internal/server"
	"github.com/CryptoGu1/books-grpc-log/internal/service"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Не найден файл .env")
	}

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- DEBUG CONFIG ---")
	fmt.Printf("URI: '%s'\n", cfg.DB.URI)
	fmt.Printf("User: '%s'\n", cfg.DB.Username)
	fmt.Println("--------------------")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()

	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	auditRepo := repo.NewAudit(db)
	auditSevice := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditSevice)
	srv := server.New(auditSrv)

	fmt.Println("Server Starter", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
