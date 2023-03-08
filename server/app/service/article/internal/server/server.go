package server

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Service interface {
	Start()
	Stop()
}

type ServerGroup struct {
	services []Service
	group    sync.WaitGroup
}

func (sg *ServerGroup) Add(service Service) {
	sg.services = append(sg.services, service)
}

func (sg *ServerGroup) Start() {
	sg.group.Add(len(sg.services))

	for i := range sg.services {
		service := sg.services[i]
		go service.Start()
	}

	go func() {
		defer sg.Stop()
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-sigs
	}()
	sg.group.Wait()
}

func (sg *ServerGroup) Stop() {
	for i := range sg.services {
		defer sg.group.Done()
		service := sg.services[i]
		service.Stop()
	}
}
