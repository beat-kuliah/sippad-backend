package main

import (
	"github/beat-kuliah/sip_pad_backend/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(8000)
}
