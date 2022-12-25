package routes

import (
	"restaurant/routes/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers())
	r.GET("/users/:user_id", controllers.GetUser())
	r.POST("/users/signup", controllers.SignUp())
	r.POST("/users/login", controllers.Login())
}
