package orders

import (
	"context"

	"github.com/google/uuid"
	"github.com/noradomi/coffeeshop-clone/internal/counter/domain"
)

type (
	OrderRepo interface {
		GetAll(context.Context) ([]*domain.Order, error)
		GetByID(context.Context, uuid.UUID) (*domain.Order, error)
		Create(context.Context, *domain.Order) error
		Update(context.Context, *domain.Order) (*domain.Order, error)
	}

	UseCase interface {
		GetListOrderFulfillment(context.Context) ([]*domain.Order, error)
		PlaceOrder(context.Context, *domain.PlaceOrderModel) error
	}
)
