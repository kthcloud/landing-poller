package workers

import (
	"landing-api-poller/pkg/app"
	"log"
)

func Start(ctx *app.Context) {
	log.Println("starting workers")

	go StatWorker()
	go CapacitiesWorker(ctx)
	go StatusWorker()
}