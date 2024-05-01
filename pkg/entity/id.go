package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(stringID string) (ID, error) {
	id, err := uuid.Parse(stringID)
	return ID(id), err
}
