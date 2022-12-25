package routes

import (
	"github.com/gin-gonic/gin"

	"restaurant/routes/controllers"
)

func NoteRoutes(r *gin.Engine) {
	r.GET("/notes", controllers.GetNotes())
	r.GET("/notes/:note_id", controllers.GetNote())
	r.POST("/notes", controllers.CreateNote())
	r.PATCH("/notes/:note_id", controllers.UpdateNote())
}
