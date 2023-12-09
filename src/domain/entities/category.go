package entities

import "time"

type CategoryEntity struct {
	ID          uint64
	UUID        string
	Name        string
	Description string
	CreatedAt   time.Time
}
