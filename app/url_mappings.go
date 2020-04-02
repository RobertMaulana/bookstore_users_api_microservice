package app

import (
	"github.com/RobertMaulana/bookstore_users_api_microservice/controllers/ping"
	"github.com/RobertMaulana/bookstore_users_api_microservice/controllers/users"

)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}