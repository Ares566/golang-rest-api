package repository

import (
	"rest-api-endpoints/domain/entity"
)

type TagsRepository interface {
	// Create part
	Create(*entity.CreateTag) error

	// Fetch Read part
	Fetch(accountID int64) (entity.Tags, error)
	FetchByProduct(accountID int64) (entity.TagsByProduct, error)

	// Update part
	Update(*entity.CreateTag) error

	// Delete part
	Delete(*entity.CreateTag) error
}
