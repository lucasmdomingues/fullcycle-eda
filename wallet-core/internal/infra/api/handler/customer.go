package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lucasmdomingues/wallet-core/internal/usecase/customer"
	"github.com/lucasmdomingues/wallet-core/pkg/web"
)

type CustomerHandler struct {
	createUsecase *customer.CreateCustomerUsecase
}

func NewCustomerHandler(createUsecase *customer.CreateCustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		createUsecase: createUsecase,
	}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var dto customer.CreateCustomerInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		log.Println("failed to decode input", err)
		web.EncodeJSON(w, http.StatusBadRequest, err)
		return
	}

	output, err := h.createUsecase.Execute(dto)
	if err != nil {
		log.Println("failed to create customer", err)
		web.EncodeJSON(w, http.StatusBadRequest, err)
		return
	}

	web.EncodeJSON(w, http.StatusCreated, output)
}
