package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/shanemalachow/stripcall/internal"
	"log"
)

func main() {
	configLocation := flag.String("config", "stripcall.conf", "Location of the StripCall config file")
	config := stripcall.ParseConfig(*configLocation)
	flag.Parse()

	dependencies := stripcall.DependencyMap{
		Conf: config,
		DB:   stripcall.Connect(config["dbType"], config["dbConnect"]),
	}
	r := mux.NewRouter()

	s := stripcall.Setup(r, dependencies)

	//Start the router, logging any fatal errors
	log.Fatal(s.ListenAndServe())
}
