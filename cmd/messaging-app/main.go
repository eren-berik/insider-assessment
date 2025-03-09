package main

import (
	"insider-assesment/pkg/api"
	"insider-assesment/pkg/app"
)

func main() {
	opt := options()
	serviceProvider := app.RegisterService(opt)
	optionProvider := app.RegisterOptions(opt)
	worker := app.NewWorker(serviceProvider, optionProvider)
	server := api.NewServer(serviceProvider, optionProvider, worker)
	server.Run()
}
