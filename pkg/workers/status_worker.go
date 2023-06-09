package workers

import (
	"context"
	"fmt"
	"landing-poller/models"
	"landing-poller/models/dto"
	"landing-poller/pkg/conf"
	"landing-poller/utils/requestutils"
	"log"
	"sync"
	"time"
)

func GetHostStatuses() ([]dto.HostStatus, error) {
	outputs := make([]*dto.HostStatus, len(conf.Hosts))

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for idx, host := range conf.Hosts {

		wg.Add(1)
		go func(idx int, host conf.Host) {
			makeError := func(err error) error {
				return fmt.Errorf("failed to get  for host %s. details: %s", host.IP.String(), err)
			}

			result, err := requestutils.DoRequest("GET", host.ApiURL("/status"), nil, nil)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			var hostStatus dto.HostStatus
			err = requestutils.ParseBody(result.Body, &hostStatus)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}
			hostStatus.Name = conf.Hosts[idx].Name

			mu.Lock()
			outputs[idx] = &hostStatus
			mu.Unlock()

			wg.Done()
		}(idx, host)
	}

	wg.Wait()

	var result []dto.HostStatus

	for _, output := range outputs {
		if output != nil {
			result = append(result, *output)
		}
	}

	return result, nil
}

func StatusWorker() {
	makeError := func(err error) error {
		return fmt.Errorf("status worker experienced an issue: %s", err)
	}

	for {
		hostsStatuses, err := GetHostStatuses()
		if err != nil {
			log.Println(makeError(err))
			time.Sleep(StatusSleep)
			continue
		}

		if len(hostsStatuses) == 0 {
			log.Println(makeError(fmt.Errorf("no hosts statuses received")))
		} else {
			status := dto.Status{
				Hosts: hostsStatuses,
			}

			statusDB := dto.StatusDB{
				Status:    status,
				Timestamp: time.Now(),
			}

			_, err = models.StatusCollection.InsertOne(context.TODO(), statusDB)
			if err != nil {
				log.Println(makeError(err))
				return
			}
		}

		time.Sleep(StatusSleep)
	}
}
