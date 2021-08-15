package repository

import "rest-api-endpoints/domain/entity"

type NotesRepository interface {
	// Create part
	Create(note *entity.CreateNote) error

	// Fetch Read part
	Fetch(accountID int64, productID int64) (entity.Notes, error)

	// Update part
	Update(note *entity.CreateNote) error

	// Delete part
	Delete(note *entity.CreateNote) error
}
