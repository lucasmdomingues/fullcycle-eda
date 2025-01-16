package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasmdomingues/wallet-balance/internal/usecase/balance"
	"github.com/lucasmdomingues/wallet-balance/pkg/web"
)

type BalanceHandler struct {
	usecase *balance.FindByAccountIDUsecase
}

func NewBalanceHandler(usecase *balance.FindByAccountIDUsecase) *BalanceHandler {
	return &BalanceHandler{usecase}
}

func (h *BalanceHandler) FindByAccountID(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountID")
	if accountID == "" {
		web.EncodeJSON(w, http.StatusBadRequest, errors.New("account id cannot be empty"))
		return
	}

	output, err := h.usecase.Execute(accountID)
	if err != nil {
		web.EncodeJSON(w, http.StatusInternalServerError, err)
		return
	}

	web.EncodeJSON(w, http.StatusOK, output)
}
