package application

import (
	"rest-api-endpoints/domain/entity"
	"rest-api-endpoints/domain/repository"
)

type TagApplication struct {
	tagsRepository repository.TagsRepository
}

func NewTagApplication(tg repository.TagsRepository) *TagApplication {
	if tg == nil {
		return nil
	}

	return &TagApplication{tg}
}

func (tg *TagApplication) Fetch(accountID int64) (entity.Tags, error) {
	return tg.tagsRepository.Fetch(accountID)
}
func (tg *TagApplication) FetchByProduct(accountID int64) (entity.TagsByProduct, error) {
	return tg.tagsRepository.FetchByProduct(accountID)
}

// Create is
func (tg *TagApplication) Create(t *entity.CreateTag) error {
	return tg.tagsRepository.Create(t)
}

// Update is
func (tg *TagApplication) Update(t *entity.CreateTag) error {
	return tg.tagsRepository.Update(t)
}

// Delete is
func (tg *TagApplication) Delete(t *entity.CreateTag) error {
	return tg.tagsRepository.Delete(t)
}
