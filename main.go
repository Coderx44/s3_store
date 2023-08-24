package main

import (
	"github.com/Coderx44/s3_store/server"
)

func main() {

	err := server.StartApp()
	if err != nil {
		panic(err)
	}
}
