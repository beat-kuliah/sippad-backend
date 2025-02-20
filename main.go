package main

import (
	"github/beat-kuliah/finbest_backend/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(3000)
}
