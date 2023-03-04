package main

import (
	"github.com/ismdeep/notification/app/server/api"
	"github.com/ismdeep/notification/app/server/worker"
)

func main() {
	go worker.Run()
	api.Run()
}
