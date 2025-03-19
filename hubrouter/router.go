package hubrouter

import (
	"github.com/Dpyde/Omchu/internal/ws"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

var r *gin.Engine

func InitRouter( wsHandler *ws.Handler) {
	r = gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},  // Allow all origins
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,            // Allow cookies
		MaxAge:           12 * time.Hour,  // Cache preflight request results for 12 hours
	}))

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
