package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lucasmdomingues/wallet-core/internal/usecase/account"
	"github.com/lucasmdomingues/wallet-core/pkg/web"
)

type AccountHandler struct {
	createUsecase *account.CreateAccountUsecase
}

func NewAccountHandler(createUsecase *account.CreateAccountUsecase) *AccountHandler {
	return &AccountHandler{
		createUsecase: createUsecase,
	}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto account.CreateAccountInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		log.Println("failed to decode input", err)
		web.EncodeJSON(w, http.StatusBadRequest, err)
		return
	}

	output, err := h.createUsecase.Execute(dto)
	if err != nil {
		log.Println("failed to create account", err)
		web.EncodeJSON(w, http.StatusBadRequest, err)
		return
	}

	web.EncodeJSON(w, http.StatusCreated, output)
}
