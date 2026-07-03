package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
)

func CORS() cors.Config {
	origins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	return cors.Config{
		AllowOrigins: origins,
		
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	}
}