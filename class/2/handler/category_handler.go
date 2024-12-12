package handler

import (
	"20241212/class/2/domain"
	"20241212/class/2/helper"
	"20241212/class/2/service"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryHandler interface {
	ShowAllCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
}

type categoryHandler struct {
	log     *zap.Logger
	service *service.Service
}

func NewCategoryHandler(log *zap.Logger, service *service.Service) CategoryHandler {
	return &categoryHandler{log, service}
}

func (ch *categoryHandler) ShowAllCategory(c *gin.Context) {
	pageStr := c.Query("page")

	page, _ := strconv.Atoi(pageStr)

	categories, err := ch.service.Category.ShowAllCategory(page)
	if err != nil {
		BadResponse(c, "Failed to retrived categories: "+err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "successfully retrived categories", http.StatusOK, categories)
}

func (ch *categoryHandler) DeleteCategory(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	if err := ch.service.Category.DeleteCategory(id); err != nil {
		BadResponse(c, "Failed to deleted categoriy: "+err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "successfully deleted category", http.StatusOK, id)

}

func (ch *categoryHandler) GetCategoryByID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	category, err := ch.service.Category.GetCategoryByID(id)
	if err != nil {
		BadResponse(c, "Failed to Retrieved categoriy: "+err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "successfully Retrieved category", http.StatusOK, category)
}

func (ch *categoryHandler) CreateCategory(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		BadResponse(c, "Invalid form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		BadResponse(c, "No image provided", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup

	responses, err := helper.Upload(&wg, []*multipart.FileHeader{files[0]})
	if err != nil || len(responses) == 0 {
		BadResponse(c, "Failed to upload image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	name := c.PostForm("name")
	if name == "" {
		BadResponse(c, "Name is required", http.StatusBadRequest)
		return
	}

	imageURL := responses[0].Data.Url

	// Buat entitas kategori baru
	category := domain.Category{
		Name: name,
		Icon: imageURL,
	}

	// Simpan kategori menggunakan service
	err = ch.service.Category.CreateCategory(&category)
	if err != nil {
		BadResponse(c, "Failed to create category: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Berikan respon sukses
	GoodResponseWithData(c, "Category created successfully", http.StatusCreated, category)
}

func (ch *categoryHandler) UpdateCategory(c *gin.Context) {

	category := domain.Category{}
	id, _ := strconv.Atoi(c.Param("id"))

	form, _ := c.MultipartForm()

	var imageURL string
	files := form.File["images"]
	if len(files) > 0 {
		var wg sync.WaitGroup
		responses, err := helper.Upload(&wg, []*multipart.FileHeader{files[0]})
		if err != nil || len(responses) == 0 {
			BadResponse(c, "Failed to upload image: "+err.Error(), http.StatusInternalServerError)
			return
		}
		imageURL = responses[0].Data.Url
	}

	name := c.PostForm("name")
	if name == "" {
		name = category.Name
	}

	category = domain.Category{
		ID:   uint(id),
		Name: name,
		Icon: imageURL,
	}

	err := ch.service.Category.UpdateCategory(id, &category)
	if err != nil {
		BadResponse(c, "Category not found: "+err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "Category updated successfully", http.StatusOK, category)
}
