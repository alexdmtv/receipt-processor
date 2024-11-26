package httpserver

import (
	"net/http"
	"receipt-processor/internal/infrastructure/httpserver/handler"
)

type Server struct {
	receiptHandler *handler.ReceiptHandler
}

func NewServer(receiptHandler *handler.ReceiptHandler) *Server {
	return &Server{
		receiptHandler: receiptHandler,
	}
}

func (s *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /receipts/process", s.receiptHandler.CreateReceipt)
	mux.HandleFunc("GET /receipts/{id}/points", s.receiptHandler.GetReceiptPoints)

	return mux
}

func (s *Server) Start(port string) error {
	mux := s.SetupRoutes()
	return http.ListenAndServe(port, mux)
}
