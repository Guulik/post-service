package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"posts/internal/app"
	"syscall"
)

func main() {

	ctx := context.Background()
	//setupLogger(cfg.Env)

	a := app.New()

	go func() {
		a.Run()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	if err := a.Stop(ctx); err != nil {
		fmt.Println(fmt.Errorf("failed to gracefully stop app err=%s", err))
	}

	fmt.Println("Gracefully stopped")
}
