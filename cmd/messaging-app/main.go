package main

import (
	"insider-assesment/pkg/app"
)

func main() {
	opt := options()
	_ = app.RegisterService(opt)

}
