package ping

import "github.com/gin-gonic/gin"

// AddRoutes godoc
// @Summary Ping endpoint
// @Description ping this server
// @Tags ping
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PingResponse
// @Router /ping [get]
func AddRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
