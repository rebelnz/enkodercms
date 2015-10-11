package main

import (
	"bitbucket.org/enkdr/enkoder/config"
	r "bitbucket.org/enkdr/enkoder/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := r.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	fmt.Printf("running on %s, %s mode\n", config.Conf().AppPort, os.Getenv("ENKDR_ENV"))
	log.Fatal(http.ListenAndServe(config.Conf().AppPort, router))
}
