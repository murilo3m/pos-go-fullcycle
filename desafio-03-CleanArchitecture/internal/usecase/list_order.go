package usecase

import (
	"20-CleanArch/internal/entity"
	"20-CleanArch/pkg/events"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersOutput []ListOrdersOutputDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	c.OrderListed.SetPayload(ordersOutput)
	c.EventDispatcher.Dispatch(c.OrderListed)

	return ordersOutput, nil
}
