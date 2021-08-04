package interfaces

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-endpoints/application"
	"rest-api-endpoints/util/response"
)


type TagsREST struct {
	TagsApp *application.TagApplication
}


func NewTagsREST(ta *application.TagApplication) *TagsREST {
	if ta == nil {
		return nil
	}

	return &TagsREST{ta}
}


func (th *TagsREST) Fetch(c *gin.Context) {

	// TODO
	// bind
	// validation

	// logic
	result, err := th.TagsApp.Fetch(c)
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", result)
}
func (th *TagsREST) FetchByProducts(c *gin.Context) {

	// TODO
	// bind
	// validation

	// logic
	result, err := th.TagsApp.FetchByProduct(c)
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", result)
}

