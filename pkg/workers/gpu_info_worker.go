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

func GetHostGpuInfo() ([]dto.HostGPUInfo, error) {

	outputs := make([]*dto.HostGPUInfo, len(conf.Hosts))

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for idx, host := range conf.Hosts {

		wg.Add(1)
		go func(idx int, host conf.Host) {
			makeError := func(err error) error {
				return fmt.Errorf("failed to get  for host %s. details: %s", host.IP.String(), err)
			}

			result, err := requestutils.DoRequest("GET", host.ApiURL("/gpuInfo"), nil, nil)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			var hostGpus []dto.HostGPU
			err = requestutils.ParseBody(result.Body, &hostGpus)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			hostGpuInfo := dto.HostGPUInfo{
				Name: conf.Hosts[idx].Name,
				GPUs: hostGpus,
			}

			mu.Lock()
			outputs[idx] = &hostGpuInfo
			mu.Unlock()

			wg.Done()
		}(idx, host)
	}

	wg.Wait()

	var result []dto.HostGPUInfo

	for _, output := range outputs {
		if output != nil {
			result = append(result, *output)
		}
	}

	return result, nil
}

func GpuInfoWorker() {
	makeError := func(err error) error {
		return fmt.Errorf("gpu info worker experienced an issue: %s", err)
	}

	for {
		hostGpuInfo, err := GetHostGpuInfo()
		if err != nil {
			log.Println(makeError(err))
			time.Sleep(GpuInfoSleep)
			continue
		}

		if len(hostGpuInfo) == 0 {
			log.Println(makeError(fmt.Errorf("no host gpu info was found")))
		} else {
			gpuInfoDB := dto.GpuInfoDB{
				GpuInfo: dto.GpuInfo{
					HostGPUInfo: hostGpuInfo,
				},
				Timestamp: time.Now().UTC(),
			}

			_, err = models.GpuInfoCollection.InsertOne(context.TODO(), gpuInfoDB)
			if err != nil {
				log.Println(makeError(err))
				continue
			}
		}

		time.Sleep(GpuInfoSleep)
	}
}
