package main

import (
	"CodingExercise/lib"
	"CodingExercise/shared/helpers"
	"CodingExercise/shared/log"
	"context"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Context to cancel functions
	ctx, cancel := context.WithCancel(context.Background())

	if len(os.Args) != 3 {
		fmt.Println("Usage: main  [path to config directory] [path to log directory]:", os.Args)
		return
	}

	configPath := os.Args[1]
	logPath := os.Args[2]

	go helpers.CaptureSignal(cancel)

	lib.ParseConfiguration(configPath, logPath)

	lib.ConnectDB(ctx)

	router := log.Logger(lib.GetRoutes())

	fmt.Println("Starting HTTP Server: ", lib.ServiceConfig.Service.Port)
	log.Debug(ctx, "Starting HTTP Server: ", lib.ServiceConfig.Service.Port)

	err := http.ListenAndServe(":"+lib.ServiceConfig.Service.Port, router)
	log.Error(ctx, " HTTP ListenAndServe err : ", err)
}
