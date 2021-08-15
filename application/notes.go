package application

import (
	"rest-api-endpoints/domain/entity"
	"rest-api-endpoints/domain/repository"
)

type NotesApplication struct {
	notesRepository repository.NotesRepository
}

func NewNotesApplication(nt repository.NotesRepository) *NotesApplication {
	if nt == nil {
		return nil
	}

	return &NotesApplication{nt}
}

// Fetch is
func (nt *NotesApplication) Fetch(accountID int64, productId int64) (entity.Notes, error) {
	return nt.notesRepository.Fetch(accountID, productId)
}

// Create is
func (nt *NotesApplication) Create(t *entity.CreateNote) error {
	return nt.notesRepository.Create(t)
}

// Update is
func (nt *NotesApplication) Update(t *entity.CreateNote) error {
	return nt.notesRepository.Update(t)
}

// Delete is
func (nt *NotesApplication) Delete(t *entity.CreateNote) error {
	return nt.notesRepository.Delete(t)
}
