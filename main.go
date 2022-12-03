package main

import (
	"SK-todo/model"
	"SK-todo/router"
)

func main() {
	model.Database("./test1.db")
	r := router.NewRouter()
	_ = r.Run(":8080")
}
