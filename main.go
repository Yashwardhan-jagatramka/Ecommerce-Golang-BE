package main

import (
	"github.com/Yashwardhan-jagatramka/ecommerce-golang-BE/controllers"
	"github.com/Yashwardhan-jagatramka/ecommerce-golang-BE/database"
	"github.com/Yashwardhan-jagatramka/ecommerce-golang-BE/middleware"
	"github.com/Yashwardhan-jagatramka/ecommerce-golang-BE/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	group := router.Group("/ecommerce")

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	group.Handle(http.MethodGet, "/addtocart", app.AddToCart())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
