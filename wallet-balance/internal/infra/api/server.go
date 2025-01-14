package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasmdomingues/wallet-balance/internal/infra/api/handlers"
	"github.com/lucasmdomingues/wallet-balance/internal/usecase/account"
)

type Server struct {
	r *chi.Mux
}

func (s Server) Start() error {
	return http.ListenAndServe(":3003", s.r)
}

func NewAPI(findAccountByIDUsecase *account.FindByIDUsecase) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	accountHandler := handlers.NewAccountHandler(findAccountByIDUsecase)

	r.Get("/balances/{accountID}", accountHandler.FindByID)

	return &Server{r}
}
