package main

import (
	"fmc-newest/internal/app"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		return
	}
	quit := make(chan os.Signal, 1)
	go func() {
		fmt.Println("SERVER START")
		log.Fatal(app.ConfigureAndRun())
	}()
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
}
