package app

import "vanthanh.com/bookstore_users_api/controllers"

func mapUrl() {
	router.GET("/ping", controllers.Pong)
	router.GET("/users", controllers.GetUser)
	router.GET("/users/:id", controllers.GetUserById)
	router.POST("/users", controllers.CreateUser)
}
