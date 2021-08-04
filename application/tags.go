package application

import (
	"github.com/gin-gonic/gin"
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

// Fetch is
func (tg *TagApplication) Fetch(c *gin.Context) (entity.Tags, error) {
	return tg.tagsRepository.Fetch(c)
}
func (tg *TagApplication) FetchByProduct(c *gin.Context) (entity.TagsByProduct, error) {
	return tg.tagsRepository.FetchByProduct(c)
}

// Create is
func (tg *TagApplication) Create(t *entity.CreateTag) error {
	return tg.tagsRepository.Create(t)
}

