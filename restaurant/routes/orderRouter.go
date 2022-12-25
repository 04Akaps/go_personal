package routes

import (
	"github.com/gin-gonic/gin"

	"restaurant/routes/controllers"
)

func OrderRoutes(r *gin.Engine) {
	r.GET("/orders", controllers.GetOrders())
	r.GET("/orders/:order_id", controllers.GetOrder())
	r.POST("/orders", controllers.CreateOrder())
	r.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
