package pingapp

import "github.com/gin-gonic/gin"

func PingRouter(g *gin.RouterGroup) {
	pingR := g.Group("/server")
	{
		pingR.GET("/ping", PingHandler)
		pingR.POST("/ping", PingHandler)
	}
}
