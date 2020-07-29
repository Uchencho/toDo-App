package main

import (
	"log"
	"net/http"

	"github.com/Uchencho/toDo-App/app"
)

func main() {

	file := app.TheLogger()

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// defer file.Close()

	http.HandleFunc("/healthcheck", app.Healthcheck)
	http.HandleFunc("/create", app.CreateEntryEndpoint)
	http.HandleFunc("/tasks", app.ListAPIView)
	app.InfoLogger.Println(app.GetServerAddress())
	if err := http.ListenAndServe(app.GetServerAddress(), nil); err != http.ErrServerClosed {
		app.ErrorLogger.Println(err)
	}
}
