package entity

import "github.com/google/uuid"

type Id = uuid.UUID

func NewId() Id {
	return Id(uuid.New())
}

func ParseId(s string) (Id, error) {
	id, err := uuid.Parse(s)

	return Id(id), err
}
