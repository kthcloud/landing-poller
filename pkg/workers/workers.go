package workers

import (
	"landing-poller/pkg/app"
	"log"
	"time"
)

var CapacitiesSleep = 60 * time.Second
var StatusSleep = 1 * time.Second
var StatsSleep = 60 * time.Second
var GpuInfoSleep = 300 * time.Second

func Start(ctx *app.Context) {
	log.Println("starting workers")

	go StatWorker()
	go CapacitiesWorker(ctx)
	go StatusWorker()
	go GpuInfoWorker()
}
