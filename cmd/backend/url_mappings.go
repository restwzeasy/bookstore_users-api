package backend

import (
	"github.com/restwzeasy/bookstore_users-api/pkg/api/health"
	"github.com/restwzeasy/bookstore_users-api/pkg/api/users"
)

func mapUrls() {
	router.GET("/health", health.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
