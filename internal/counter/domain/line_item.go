package domain

import (
	"github.com/google/uuid"
	shared "github.com/noradomi/coffeeshop-clone/internal/pkg/shared_kernel"
)

type LineItem struct {
	ID             uuid.UUID
	ItemType       shared.ItemType
	Name           string
	Price          float32
	ItemStatus     shared.Status
	IsBaristaOrder bool
	OrderID        uuid.UUID // shadow field ?
}

func NewLineItem(itemType shared.ItemType, name string, price float32, itemStatus shared.Status, isBaristaOrder bool) *LineItem {
	return &LineItem{
		ID:             uuid.New(),
		ItemType:       itemType,
		Name:           name,
		Price:          price,
		ItemStatus:     itemStatus,
		IsBaristaOrder: isBaristaOrder,
	}
}
