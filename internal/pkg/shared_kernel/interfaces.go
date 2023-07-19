package shared_kernel

import "time"

type (
	DomainEvent interface {
		CreatedAt() time.Time
		Identity() string
	}
)
