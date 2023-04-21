package main

import (
	"landing-poller/models"
	"landing-poller/pkg/app"
	"landing-poller/pkg/conf"
	"landing-poller/pkg/workers"
)

func main() {
	context := app.Context{}

	conf.Setup()
	models.Setup()
	workers.Start(&context)
	select {}
}
