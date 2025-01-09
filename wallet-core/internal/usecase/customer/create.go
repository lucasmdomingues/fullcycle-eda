package customer

import (
	"log"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
	"github.com/lucasmdomingues/wallet-core/internal/domain/repository"
)

type CreateCustomerUsecase struct {
	repository repository.Customer
}

func NewCreateCustomerUsecase(repository repository.Customer) *CreateCustomerUsecase {
	return &CreateCustomerUsecase{
		repository: repository,
	}
}

func (usecase *CreateCustomerUsecase) Execute(input CreateCustomerInputDTO) (CreateCustomerOutput, error) {
	customer, err := entity.NewCustomer(input.Name, input.Email)
	if err != nil {
		return CreateCustomerOutput{}, err
	}

	err = usecase.repository.Save(customer)
	if err != nil {
		log.Println("failed to save customer", err)
		return CreateCustomerOutput{}, err
	}

	return CreateCustomerOutput{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}, nil
}
