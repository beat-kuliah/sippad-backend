package main

import (
	"github/beatfraps/finbest_backend/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(3000)
}
