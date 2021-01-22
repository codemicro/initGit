package main

import (
	"github.com/codemicro/initGit/internal/app"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(0)
	}()
	app.Run()
}
