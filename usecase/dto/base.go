package dto

import "github.com/google/uuid"

type CreateAreainfo struct {
	ID          uuid.UUID
	Name        string
	Description string
	Code        int64
}

type UpdateAreainfo struct {
	ID          uuid.UUID
	Name        string
	Description string
	Code        int64
}
type AreaInfo struct {
	ID          uuid.UUID
	Name        string
	Description string
	Code        int64
	//Tags        []Tag
}
