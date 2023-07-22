package main

import (
	api "auth0-example"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := api.Router()
	router.Run(":" + port) //nolint:errcheck
}
