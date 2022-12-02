package main

import (
	"SK-todo/model"
	"SK-todo/router"
)

func main() {
	model.Database("file::memory:?cache=shared")
	r := router.NewRouter()
	_ = r.Run(":8080")
}
