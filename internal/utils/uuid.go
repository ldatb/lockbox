package utils

import "github.com/google/uuid"

// NewUUID create a new entity ID
func NewUUID() uuid.UUID {
	return uuid.New()
}

// StringToUUID convert a string to an UUID
func StringToUUID(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	return id, err
}
