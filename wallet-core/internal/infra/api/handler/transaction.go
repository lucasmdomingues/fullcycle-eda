package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lucasmdomingues/wallet-core/internal/usecase/transaction"
	"github.com/lucasmdomingues/wallet-core/pkg/web"
)

type TransactionHandler struct {
	createUsecase *transaction.CreateTransactionUsecase
}

func NewTransactionHandler(createUsecase *transaction.CreateTransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		createUsecase: createUsecase,
	}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto transaction.CreateTransactionInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		log.Println("failed to decode input", err)
		web.EncodeJSON(w, http.StatusBadRequest, err)
		return
	}

	output, err := h.createUsecase.Execute(r.Context(), dto)
	if err != nil {
		log.Println("failed to create transaction", err)
		web.EncodeJSON(w, http.StatusBadRequest, err)
		return
	}

	web.EncodeJSON(w, http.StatusCreated, output)
}
