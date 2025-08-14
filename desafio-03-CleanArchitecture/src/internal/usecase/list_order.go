package usecase

import "cleanArch/internal/entity"

type ListOrdersOutputDTO struct {
	List  []*entity.Order `json:"orders"`
	Total int             `json:"total"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}
	dto := ListOrdersOutputDTO{
		List:  orders,
		Total: len(orders),
	}
	return dto, nil

}
