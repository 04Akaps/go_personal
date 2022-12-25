package main

import (
	"os"
	"restaurant/database"
	"restaurant/middleware"
	"restaurant/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "collectionName")

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()

	router.User(gin.Logger())
	router.Use(middleware.Authentication())

	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.NoteRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)

}
