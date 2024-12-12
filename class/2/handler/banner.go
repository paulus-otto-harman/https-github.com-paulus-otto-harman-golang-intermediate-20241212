package handler

import (
	"20241212/class/2/domain"
	"20241212/class/2/helper"
	"20241212/class/2/service"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ControllerBanner struct {
	service service.ServiceBanner
	logger  *zap.Logger
}

func NewControllerBanner(service service.ServiceBanner, logger *zap.Logger) *ControllerBanner {
	return &ControllerBanner{service: service, logger: logger}
}

// @Summary Get All Banner
// @Description Endpoint Fetch All Banner
// @Tags GetAllBanner
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response{data=[]domain.Banner} "Get All Success"
// @Failure 500 {object} handler.Response "server error"
// @Router  /banner [get]
func (ctrl *ControllerBanner) GetAll(c *gin.Context) {
	banners, err := ctrl.service.GetAll()
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "Get Banners success", http.StatusOK, banners)
}
func (ctrl *ControllerBanner) GetById(c *gin.Context) {
	id, err := helper.Uint(c.Param("id"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	banner, err := ctrl.service.GetById(id)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	GoodResponseWithData(c, "Get Banner success", http.StatusOK, banner)
}
func (ctrl *ControllerBanner) Create(c *gin.Context) {
	doUpload := true
	var respThirdParty []domain.CdnResponse
	_, err := c.FormFile("images")
	if err != nil {
		doUpload = false
		respThirdParty = append(respThirdParty, domain.CdnResponse{Data: struct {
			FileId      string "json:\"fileId\""
			Name        string "json:\"name\""
			Size        int    "json:\"size\""
			VersionInfo struct {
				Id   string "json:\"id\""
				Name string "json:\"name\""
			} "json:\"versionInfo\""
			FilePath     string      "json:\"filePath\""
			Url          string      "json:\"url\""
			FileType     string      "json:\"fileType\""
			Height       int         "json:\"height\""
			Width        int         "json:\"width\""
			ThumbnailUrl string      "json:\"thumbnailUrl\""
			AITags       interface{} "json:\"AITags\""
		}{Url: ""}})
	}
	if doUpload {
		form, err := c.MultipartForm()
		if err != nil {
			log.Println("Error reading form data:", err)
			BadResponse(c, "Invalid form data: "+err.Error(), http.StatusBadRequest)
			return
		}
		files := form.File["images"]
		for _, file := range files {
			log.Println("File size:", file.Size)
		}

		var wg sync.WaitGroup
		respThirdParty, err = helper.Upload(&wg, files)
		if err != nil {
			BadResponse(c, "Failed to upload images: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	banner := domain.Banner{
		Title:     c.PostForm("title"),
		PathPage:  c.PostForm("pathPage"),
		StartDate: c.PostForm("startDate"),
		EndDate:   c.PostForm("endDate"),
		IsPublish: false,
		ImageUrl:  respThirdParty[0].Data.Url,
	}
	err = ctrl.service.Create(&banner)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	GoodResponseWithData(c, "This Banner was successfully added", http.StatusCreated, "banner")
}
func (ctrl *ControllerBanner) Edit(c *gin.Context) {
	id, err := helper.Uint(c.Param("id"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	doUpload := true
	var respThirdParty []domain.CdnResponse
	_, err = c.FormFile("images")
	if err != nil {
		doUpload = false
		respThirdParty = append(respThirdParty, domain.CdnResponse{Data: struct {
			FileId      string "json:\"fileId\""
			Name        string "json:\"name\""
			Size        int    "json:\"size\""
			VersionInfo struct {
				Id   string "json:\"id\""
				Name string "json:\"name\""
			} "json:\"versionInfo\""
			FilePath     string      "json:\"filePath\""
			Url          string      "json:\"url\""
			FileType     string      "json:\"fileType\""
			Height       int         "json:\"height\""
			Width        int         "json:\"width\""
			ThumbnailUrl string      "json:\"thumbnailUrl\""
			AITags       interface{} "json:\"AITags\""
		}{Url: ""}})
	}
	if doUpload {
		form, err := c.MultipartForm()
		if err != nil {
			log.Println("Error reading form data:", err)
			BadResponse(c, "Invalid form data: "+err.Error(), http.StatusBadRequest)
			return
		}
		files := form.File["images"]
		for _, file := range files {
			log.Println("File size:", file.Size)
		}

		var wg sync.WaitGroup
		respThirdParty, err = helper.Upload(&wg, files)
		if err != nil {
			BadResponse(c, "Failed to upload images: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	banner := domain.Banner{
		ID:        id,
		Title:     c.PostForm("title"),
		PathPage:  c.PostForm("pathPage"),
		StartDate: c.PostForm("startDate"),
		EndDate:   c.PostForm("endDate"),
		IsPublish: strings.ToLower(c.PostForm("isPublish")) == "true",
		ImageUrl:  respThirdParty[0].Data.Url,
	}
	err = ctrl.service.Edit(&banner)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	GoodResponseWithData(c, "This Banner was successfully updated", http.StatusCreated, banner)
}
func (ctrl *ControllerBanner) Delete(c *gin.Context) {
	id, err := helper.Uint(c.Param("id"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	var banner domain.Banner
	banner.ID = id
	err = ctrl.service.Delete(&banner)
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "This Banner was successfully deleted", http.StatusCreated, banner)
}
