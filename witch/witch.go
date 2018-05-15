package witch

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/barryz/rmqmonitor/g"
	"github.com/barryz/rmqmonitor/witch/system"
)

const (
	commandBuildIn    = "buildin"
	commandSupervisor = "supervisor"
	commandSystemd    = "systemd"
)

func handleSignals(exitFunc func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-sigs
	log.Printf("Signal %v captured", sig)
	exitFunc()
}

func createSystem(cfg *g.GlobalConfig) system.System {
	switch cfg.Witch.Control {
	case commandBuildIn:
		return system.NewLauncher(cfg.Witch.PidFile, cfg.Witch.Command)
	case commandSupervisor:
		return system.NewSupervisor(cfg.Witch.Service)
	case commandSystemd:
		return system.NewSystemd(cfg.Witch.Service)
	}
	log.Fatalf("Invalid control '%s'", cfg.Witch.Control)
	return nil
}

func createStats() system.Stats {
	return system.NewStatsDBCtl()
}

// Launch launch witch service
func Launch() {
	cfg := g.Config()
	stats := createStats()
	sys := createSystem(cfg)
	sys.Start()

	srv := NewServer(cfg.Witch.ListenAddr, &system.SysController{System: sys}, &system.StatsController{Stats: stats}, cfg)
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("[FATAL] Start system server failed: %v", err)
		}
	}()

	handleSignals(func() {
		sys.Stop()
	})
}
