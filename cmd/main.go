package main

import (
	"context"
	"errors"
	"github.com/samber/do"
	"gocourse15/app"
	"gocourse15/pkg/myhttp"
	"gocourse15/pkg/scheduler"
	"gocourse15/providers"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// create injector
	injector := do.DefaultInjector
	ctx := context.Background()

	// listen to os interrupt signals and close the context
	ctx, cancelFunc := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer cancelFunc()

	// inject the signal notify context
	do.ProvideValue(injector, ctx)

	providers.Provide(injector)

	waitForTheEnd := &sync.WaitGroup{}

	// start the http server
	go func() {
		waitForTheEnd.Add(1)
		defer waitForTheEnd.Done()

		log.Println("Starting HTTP server")
		httpServer := do.MustInvoke[*myhttp.HTTP](injector)
		go func() {
			<-ctx.Done()
			if err := httpServer.Shutdown(); err != nil {
				log.Fatalln(err)
			}
		}()
		if err := httpServer.Start(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalln(err)
			}
		}
	}()

	// start the scheduler service
	go func() {
		waitForTheEnd.Add(1)
		defer waitForTheEnd.Done()

		log.Println("Starting scheduler")
		tasksScheduler := do.MustInvoke[*scheduler.Scheduler](injector)
		go func() {
			<-ctx.Done()
			tasksScheduler.Shutdown()
		}()
		if err := tasksScheduler.Manage(ctx); err != nil {
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()

	// test
	service := do.MustInvoke[*app.Service](injector)
	if err := service.Do(ctx); err != nil {
		log.Fatalln(err)
	}

	<-ctx.Done()
	waitForTheEnd.Wait()

	return
}
