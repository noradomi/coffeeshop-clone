package domain

import (
	"github.com/google/uuid"
	shared "github.com/noradomi/coffeeshop-clone/internal/pkg/shared_kernel"
	"time"
)

type PlaceOrderModel struct {
	CommandType     shared.CommandType
	OrderSource     shared.OrderSource
	Location        shared.Location
	LoyaltyMemberID uuid.UUID
	BaristaItems    []*OrderItemModel
	KitchenItems    []*OrderItemModel
	Timestamp       time.Time
}

type OrderItemModel struct {
	ItemType shared.ItemType
}

type ItemModel struct {
	ItemType shared.ItemType
	Price    float64
}
