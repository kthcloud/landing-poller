package workers

import (
	"landing-poller/pkg/app"
	"log"
)

func Start(ctx *app.Context) {
	log.Println("starting workers")

	go StatWorker()
	go CapacitiesWorker(ctx)
	go StatusWorker()
}
