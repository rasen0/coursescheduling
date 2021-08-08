package main

import (
	"context"
	"coursesheduling/database"
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
	err = database.InitDB(configure.DBInfo)
	if err != nil{
		log.Errorf("init database fail. %v",err)
		return
	}

	log.Print("start course scheduling system")
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
