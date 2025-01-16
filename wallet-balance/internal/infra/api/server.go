package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasmdomingues/wallet-balance/internal/infra/api/handlers"
	"github.com/lucasmdomingues/wallet-balance/internal/usecase/balance"
)

type Server struct {
	r *chi.Mux
}

func (s Server) Start() error {
	return http.ListenAndServe(":3003", s.r)
}

func NewAPI(findBalanceByAccountID *balance.FindByAccountIDUsecase) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	balanceHandler := handlers.NewBalanceHandler(findBalanceByAccountID)

	r.Get("/balances/{accountID}", balanceHandler.FindByAccountID)

	return &Server{r}
}
