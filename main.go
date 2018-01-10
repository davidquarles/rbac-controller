package main

import (
	"github.com/davidquarles/rbac-controller/pkg/controller"
	"github.com/davidquarles/rbac-controller/pkg/handlers"
)

func main() {
	eventHandler := &handlers.Default{}
	controller.Start(eventHandler)
}
