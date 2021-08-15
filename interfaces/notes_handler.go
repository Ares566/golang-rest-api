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

type NotesREST struct {
	NotesApp *application.NotesApplication
}

func NewNotesREST(app *application.NotesApplication) *NotesREST {
	if app == nil {
		return nil
	}

	return &NotesREST{app}
}

// Create
// @Summary Создание заметки
// @Description Создание заметки
// @Tags Notes
// @Accept  json
// @Produce  json
// @Param data body entity.CreateNote true "Тело запроса"
// @Success 200 {object} response.SuccessResponse
// @Failure 400,404 {object} map[string][]string
// @Router /note [post]
func (hndl *NotesREST) Create(c *gin.Context) {

	// bind
	var request entity.CreateNote
	errBadRqst := errors.New("bad request")
	if c.BindJSON(&request) == nil {
		if request.Text == "" || request.AccountID == 0 || request.ProductID == 0 {
			response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
			return
		}
	} else {
		response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
		return
	}

	// logic
	err := hndl.NotesApp.Create(&request)
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", "")
}

// Fetch godoc
// @Summary Возвращает Заметки
// @Description Возвращает все Заметки аккаунта или заметку продукта
// @Tags Notes
// @Accept  json
// @Produce  json
// @Param accountid query int true "Account ID"
// @Param productid query int false "Product aka Nm ID"
// @Success 200 {object} entity.Tags
// @Failure 400,404 {object} map[string][]string
// @Router /notes/ [get]
func (hndl *NotesREST) Fetch(c *gin.Context) {

	accId, _ := strconv.Atoi(c.Query("accountid"))
	productId, _ := strconv.Atoi(c.Query("productid"))
	result, err := hndl.NotesApp.Fetch(int64(accId), int64(productId))
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", result)
}

// Update
// @Summary Обновление Заметки
// @Description Обновление Заметки
// @Tags Notes
// @Accept  json
// @Produce  json
// @Param data body entity.CreateNote true "Тело запроса"
// @Success 200 {object} response.SuccessResponse
// @Failure 400,404 {object} map[string][]string
// @Router /note [put]
func (hndl *NotesREST) Update(c *gin.Context) {

	// bind
	var request entity.CreateNote
	errBadRqst := errors.New("bad request")
	if c.BindJSON(&request) == nil {
		if request.ProductID == 0 || request.Text == "" || request.AccountID == 0 {
			response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
			return
		}
	} else {
		response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
		return
	}
	//TODO validation

	// logic
	err := hndl.NotesApp.Update(&request)
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", "")
}

// Delete
// @Summary  Удаление заметки
// @Description Удаление заметки
// @Tags Notes
// @Accept  json
// @Produce  json
// @Param data body entity.CreateNote true "Тело запроса"
// @Success 200 {object} response.SuccessResponse
// @Failure 400,404 {object} map[string][]string
// @Router /note [delete]
func (hndl *NotesREST) Delete(c *gin.Context) {

	// bind
	var request entity.CreateNote
	errBadRqst := errors.New("bad request")
	if c.BindJSON(&request) == nil {
		if request.ProductID == 0 || request.AccountID == 0 {
			response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
			return
		}
	} else {
		response.ErrorNotFound(c, errBadRqst.Error(), errBadRqst)
		return
	}
	// logic
	err := hndl.NotesApp.Delete(&request)
	if err != nil {
		response.ErrorNotFound(c, err.Error(), err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", "")
}
