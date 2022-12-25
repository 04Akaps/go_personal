package routes

import (
	"github.com/gin-gonic/gin"

	"restaurant/routes/controllers"
)

func OrderItemRoutes(r *gin.Engine) {
	r.GET("/orderItems", controllers.GetOrderItems())
	r.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	r.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	r.POST("/orderItems", controllers.CreateOrderItem())
	r.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
