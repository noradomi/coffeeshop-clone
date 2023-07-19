package events

import (
	"context"
	"github.com/noradomi/coffeeshop-clone/internal/pkg/event"
)

type (
	BaristaOrderUpdatedEventHandler interface {
		Handle(context.Context, *event.BaristaOrderUpdated) error
	}

	KitchenOrderUpdatedEventHandler interface {
		Handle(context.Context, *event.KitchenOrderUpdated) error
	}
)
