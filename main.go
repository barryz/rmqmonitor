package main

import (
	"flag"
	"rmqmon/g"
	"time"
	"rmqmon/falcon"
)

func Collect() {
	go collect(g.Config().Interval)
}

func collect(sec int64) {
	t := time.NewTicker(time.Second * time.Duration(sec)).C
	for {
		<-t
		falcon.Collector()
	}
}

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")

	flag.Parse()

	g.ParseConfig(*cfg)

	Collect()

	select {}

}
