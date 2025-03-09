package main

import (
	"insider-assessment/pkg/api"
	"insider-assessment/pkg/app"
)

func main() {
	opt := options()
	serviceProvider := app.RegisterService(opt)
	optionProvider := app.RegisterOptions(opt)
	worker := app.NewWorker(serviceProvider, optionProvider)
	server := api.NewServer(serviceProvider, optionProvider, worker)
	server.Run()
}
