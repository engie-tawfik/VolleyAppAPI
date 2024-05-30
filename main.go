package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"volleyapp/config"
	"volleyapp/infrastructure/routes"
)

func main() {
	config.LoadConfig()
	config.ConnectToDB()
	config.InitServer()
	routes.InitRoutes()

	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("See you in next game!")
}
