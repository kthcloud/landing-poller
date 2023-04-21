package main

import (
	"landing-api-poller/models"
	"landing-api-poller/pkg/app"
	"landing-api-poller/pkg/conf"
	"landing-api-poller/pkg/workers"
)

func main() {
	context := app.Context{}

	conf.Setup()
	models.Setup()
	workers.Start(&context)
	select {}
}
