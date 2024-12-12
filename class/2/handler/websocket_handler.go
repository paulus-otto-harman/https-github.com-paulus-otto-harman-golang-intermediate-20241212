package handler

import (
	"20241212/class/2/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type WebsocketController struct {
	service  service.Service
	logger   *zap.Logger
	upgrader websocket.Upgrader
}

func NewWebsocketController(service service.Service, logger *zap.Logger) *WebsocketController {
	upgrade := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		// Izinkan semua origin (ubah sesuai kebutuhan)
		return true
	}}
	return &WebsocketController{service: service, logger: logger, upgrader: upgrade}
}

func (ctrl *WebsocketController) Listen(c *gin.Context) {
	conn, err := ctrl.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()
	fmt.Println("WebSocket :: Client connected")

	for {
		total, err := ctrl.service.Order.Summary()
		fmt.Println(int(total))
		tes, _ := json.Marshal(struct {
			SummaryDate time.Time `json:"order_date"`
			Total       int       `json:"total"`
		}{
			SummaryDate: time.Now(),
			Total:       int(total),
		})
		err = conn.WriteMessage(websocket.TextMessage, tes)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}

		time.Sleep(2 * time.Second)
	}
}
