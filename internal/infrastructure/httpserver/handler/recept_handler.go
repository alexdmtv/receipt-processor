package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"receipt-processor/internal/domain/receipt"
)

type ReceiptHandler struct {
	receiptService *receipt.Service
}

func NewReceiptHandler(receiptService *receipt.Service) *ReceiptHandler {
	return &ReceiptHandler{
		receiptService: receiptService,
	}
}

func (h *ReceiptHandler) CreateReceipt(w http.ResponseWriter, r *http.Request) {
	var receiptDTO receipt.CreateReceiptDTO
	if err := json.NewDecoder(r.Body).Decode(&receiptDTO); err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := receiptDTO.Validate(); err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	newReceipt, err := h.receiptService.Create(r.Context(), receiptDTO)
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": newReceipt.Id.String()}
	writeJSON(w, response, http.StatusCreated)
}

func (h *ReceiptHandler) GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeJSONError(w, errors.New("receipt ID is required"), http.StatusBadRequest)
		return
	}

	points, err := h.receiptService.GetReceiptPoints(r.Context(), id)
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]int64{"points": points}
	writeJSON(w, response, http.StatusOK)
}
