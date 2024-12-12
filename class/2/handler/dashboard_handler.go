package handler

import (
	"20241212/class/2/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DashboardHandler interface {
	GetEerningProduct(c *gin.Context)
	GetSummary(c *gin.Context)
	GetBestSeller(c *gin.Context)
	GetMonthlyRevenue(c *gin.Context)
}

type dashboardHandler struct {
	service *service.Service
	log     *zap.Logger
}

func NewDashboardHandler(service *service.Service, log *zap.Logger) DashboardHandler {
	return &dashboardHandler{service, log}
}

func (dh *dashboardHandler) GetEerningProduct(c *gin.Context) {

	totalEarning, err := dh.service.Dashboard.GetEarningProduct()
	if err != nil {
		BadResponse(c, "There is no earning yet", http.StatusBadRequest)
		return
	}

	data := make(map[string]interface{})

	data["total_earning"] = totalEarning

	GoodResponseWithData(c, "successfully retrieved earning", http.StatusOK, data)
}

func (dh *dashboardHandler) GetSummary(c *gin.Context) {

	summary, err := dh.service.Dashboard.GetSummary()
	if err != nil {
		BadResponse(c, "There is no summary yet: "+err.Error(), http.StatusBadRequest)
		return
	}

	GoodResponseWithData(c, "successfully retrieved earning", http.StatusOK, summary)
}

func (dh *dashboardHandler) GetBestSeller(c *gin.Context) {

	bestSellers, err := dh.service.Dashboard.GetBestSeller()
	if err != nil {
		BadResponse(c, "not found best seller: "+err.Error(), http.StatusBadRequest)
		return
	}

	GoodResponseWithData(c, "successfully retrieved", http.StatusOK, bestSellers)
}

func (dh *dashboardHandler) GetMonthlyRevenue(c *gin.Context) {

	revenue, err := dh.service.Dashboard.GetMonthlyRevenue()
	if err != nil {
		BadResponse(c, "Not Found revenue: "+err.Error(), http.StatusBadRequest)
		return
	}

	GoodResponseWithData(c, "successfully retrieved", http.StatusOK, revenue)
}
