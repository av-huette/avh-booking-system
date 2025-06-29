package main

import (
	"internal/handlers"
)

func main() {
	r := handlers.SetupRouter()
	r.Run(":8080")
}
