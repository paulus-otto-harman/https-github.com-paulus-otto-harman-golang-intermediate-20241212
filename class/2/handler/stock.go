package handler

import (
	"20241212/class/2/domain"
	"20241212/class/2/helper"
	"20241212/class/2/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ControllerStock struct {
	service service.ServiceStock
	log     *zap.Logger
}

func NewServiceStock(service service.ServiceStock, log *zap.Logger) *ControllerStock {
	return &ControllerStock{service, log}
}

func (ctrl *ControllerStock) GetDetails(c *gin.Context) {
	id, err := helper.Uint(c.Param("productVariantId"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	data, err := ctrl.service.GetDetails(int(id))
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get Detail Stock success", http.StatusOK, data)
}
func (ctrl *ControllerStock) Edit(c *gin.Context) {
	id, err := helper.Uint(c.Param("productVariantId"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	newStock := struct{ NewStock int }{}
	if err := c.ShouldBindJSON(&newStock); err != nil {
		BadResponse(c, "Bad Request (Body)", http.StatusBadRequest)
		return
	}
	fmt.Println(c.PostForm("newStock"), "<<<<<<<")
	// newStock, err := helper.Uint(c.PostForm("newStock"))
	// if err != nil {
	// 	BadResponse(c, "Bad Request (Body)", http.StatusBadRequest)
	// 	return
	// }
	data, err := ctrl.service.Edit(int(id), newStock.NewStock)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Edit Stock success", http.StatusOK, data)
}
func (ctrl *ControllerStock) Delete(c *gin.Context) {
	id, err := helper.Uint(c.Param("id"))
	if err != nil {
		BadResponse(c, "Bad Request (Params)", http.StatusBadRequest)
		return
	}
	var data domain.Stock
	data.ID = id
	if err := ctrl.service.Delete(&data); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Delete History Stock success", http.StatusOK, data)
}
