package app

import "vanthanh.com/bookstore_users_api/controllers"

func mapUrl() {
	router.GET("/ping", controllers.Pong)

	router.GET("/users", controllers.GetUser)
	router.GET("/users/:id", controllers.GetUserById)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.PATCH("/users/:id", controllers.UpdateUser)
	router.POST("/users", controllers.CreateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.GET("/users/search", controllers.FindByStatus)
}
