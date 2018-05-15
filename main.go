package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/barryz/rmqmonitor/cron"
	"github.com/barryz/rmqmonitor/falcon"
	"github.com/barryz/rmqmonitor/g"
	v "github.com/barryz/rmqmonitor/version"
	"github.com/barryz/rmqmonitor/witch"
)

func collect() {
	go metricCollector(g.Config().Interval)
}

func rotateQLog() {
	go cron.Start()
}

func witchLaunch() {
	go witch.Launch()
}

func metricCollector(sec int64) {
	t := time.NewTicker(time.Second * time.Duration(sec)).C
	for {
		<-t
		falcon.Collector()
	}
}

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	ver := flag.Bool("v", false, "show agent version")

	flag.Parse()

	if *ver {
		fmt.Println(v.Build)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	if g.Config().Enabled.Collect {
		collect()
	}

	if g.Config().Enabled.LogRotate {
		rotateQLog()
	}

	if g.Config().Enabled.Witch {
		witchLaunch()
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-signals
		log.Printf("Signal %v captured", sig)
		os.Exit(0)
	}()

	select {}
}
