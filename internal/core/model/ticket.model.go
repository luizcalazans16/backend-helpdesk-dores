package model

import "github.com/google/uuid"

type Ticket struct {
	Id          uuid.UUID
	Title       string
	Description string
	Priority    int
	Category    *Category
}
