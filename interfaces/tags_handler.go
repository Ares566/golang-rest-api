package interfaces

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-endpoints/application"
	"rest-api-endpoints/domain/entity"
	"rest-api-endpoints/interfaces/response"
	"strconv"
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

// Fetch godoc
// @Summary Отображает все теги аккаунта
// @Description Tags by Account ID
// @Tags Tags
// @Accept  json
// @Produce  json
// @Param accountid query int true "Account ID"
// @Success 200 {object} entity.Tags
// @Failure 400,404 {object} map[string][]string
// @Router /alltags/ [get]
func (th *TagsREST) Fetch(c *gin.Context) {

	accId, _ := strconv.Atoi(c.Query("accountid"))
	result, err := th.TagsApp.Fetch(int64(accId))
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", result)
}

// FetchByProducts godoc
// @Summary Возвращает все тегированные продукты аккаунта с соответствующими тегами
// @Description Tags by Account ID
// @Tags Tags
// @Accept  json
// @Produce  json
// @Param accountid query int true "Account ID"
// @Success 200 {object} entity.TagsByProduct
// @Failure 400,404 {object} map[string][]string
// @Router /tagsbyproducts/ [get]
func (th *TagsREST) FetchByProducts(c *gin.Context) {

	accId, _ := strconv.Atoi(c.Query("accountid"))
	result, err := th.TagsApp.FetchByProduct(int64(accId))
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", result)
}

// Update
// @Summary Обновление тега
// @Description Обновление тега
// @Tags Tags
// @Accept  json
// @Produce  json
// @Param data body entity.CreateTag true "Тело запроса"
// @Success 200 {object} response.SuccessResponse
// @Failure 400,404 {object} map[string][]string
// @Router /tag [put]
func (th *TagsREST) Update(c *gin.Context) {

	// bind
	var request entity.CreateTag
	errBadRqst := errors.New("bad request")
	if c.BindJSON(&request) == nil {
		if request.ID == 0 || request.TagName == "" || request.AccountID == 0 {
			response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
			return
		}
	} else {
		response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
		return
	}
	//TODO validation

	// logic
	err := th.TagsApp.Update(&request)
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", "")
}

// Create
// @Summary  Создание тега
// @Description Создание тега
// @Tags Tags
// @Accept  json
// @Produce  json
// @Param data body entity.CreateTag true "Тело запроса"
// @Success 200 {object} response.SuccessResponse
// @Failure 400,404 {object} map[string][]string
// @Router /tag [post]
func (th *TagsREST) Create(c *gin.Context) {

	// bind
	var request entity.CreateTag
	errBadRqst := errors.New("bad request")
	if c.BindJSON(&request) == nil {
		if request.TagName == "" || request.AccountID == 0 {
			response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
			return
		}
	} else {
		response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
		return
	}
	//TODO validation

	// logic
	err := th.TagsApp.Create(&request)
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", "")
}

// Delete
// @Summary  Удаление тега
// @Description Удаление тега
// @Tags Tags
// @Accept  json
// @Produce  json
// @Param data body entity.CreateTag true "Тело запроса"
// @Success 200 {object} response.SuccessResponse
// @Failure 400,404 {object} map[string][]string
// @Router /tag [delete]
func (th *TagsREST) Delete(c *gin.Context) {

	// bind
	var request entity.CreateTag
	errBadRqst := errors.New("bad request")
	if c.BindJSON(&request) == nil {
		if request.ID == 0 || request.AccountID == 0 {
			response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
			return
		}
	} else {
		response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
		return
	}
	//TODO validation

	// logic
	err := th.TagsApp.Delete(&request)
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", "")
}
