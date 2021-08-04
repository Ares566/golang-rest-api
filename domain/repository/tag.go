package repository

import (
	"github.com/gin-gonic/gin"
	"rest-api-endpoints/domain/entity"
)

type TagsRepository interface {
	Fetch(c *gin.Context) (entity.Tags, error)
	Create(*entity.CreateTag) error

	FetchByProduct(c *gin.Context) (entity.TagsByProduct, error)
}
