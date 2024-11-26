package main

import (
	"log"
	"os"
	"receipt-processor/internal/domain/receipt"
	"receipt-processor/internal/infrastructure/database/memdb"
	"receipt-processor/internal/infrastructure/database/memdb/repository"
	"receipt-processor/internal/infrastructure/httpserver"
	"receipt-processor/internal/infrastructure/httpserver/handler"
)

func main() {
	db := memdb.New()
	receiptRepo := repository.NewReceiptRepository(db)
	receiptService := receipt.NewService(receiptRepo)
	receiptHandler := handler.NewReceiptHandler(receiptService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	server := httpserver.NewServer(receiptHandler)

	log.Printf("Server starting on %s...", port)
	log.Fatal(server.Start(":" + port))
}
