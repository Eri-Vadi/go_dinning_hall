package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Eri-Vadi/go_dining_hall/application"
	"github.com/Eri-Vadi/go_dining_hall/internal/infrastructure/config"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	config.LoadConfig()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	app := application.CreateApp()
	go app.Start()

	<-sigChan

	app.Shutdown(ctx)
	cancel()
}
