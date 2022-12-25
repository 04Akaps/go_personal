package routes

import (
	"restaurant/routes/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(r *gin.Engine) {
	r.GET("/foods", controllers.GetFoods())
	r.GET("/foods/:food_id", controllers.GetFood())
	r.POST("/foods", controllers.CreateFood())
	r.PATCH("/foods/:food_id", controllers.UpdateFood())
}
