package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("Starting HTTP service at", port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Fatalln("An error occurred starting HTTP listener at port", port)
		log.Fatalln("Error: " + err.Error())
	}
}
