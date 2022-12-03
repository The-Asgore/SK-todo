package main

import (
	"SK-todo/model"
	"SK-todo/router"
)

func main() {
	model.Database("./test.db")
	r := router.NewRouter()
	_ = r.Run(":8080")
}
