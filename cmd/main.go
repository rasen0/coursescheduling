package main

import (
	"context"
	"coursesheduling/lib/config"
	"coursesheduling/lib/log"
	"coursesheduling/server"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.InitLog()
	configure, err := config.InitConfigure()
	if err != nil{
		log.Error(err)
		return
	}
	log.Print("configure:",configure)

	svr := server.NewServer(&configure)
	svr.Serve()

	notifyContext, stop := signal.NotifyContext(context.Background(), os.Kill)
	<- notifyContext.Done()
	stop()
	log.Print("shutting down gracefully, press Ctrl+C again to force")
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	if err = svr.Shutdown(timeout); err != nil{
		log.Error("shutdown fail.",err)
	}
}
