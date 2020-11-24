package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/columbio/api-gate/asbi"
	"github.com/columbio/api-gate/config"
	"github.com/columbio/api-gate/loggers"
	"github.com/columbio/api-gate/whois"

	"github.com/gorilla/mux"
)

// main - application entry point
func main() {
	// make loggers initialization
	loggers.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	loggers.Info.Printf("Application starts")
	// read config
	config.ReadConfig()
	// web app start
	handleRequests()
}

// emptyResponce - http get email
func emptyResponce(w http.ResponseWriter, r *http.Request) {
	// Print nothing in HTTP Responce
	loggers.Info.Printf("emptyResponce(): Print empty HTTP Response")
	fmt.Fprintf(w, "")
}

// handleRequests - method to handle requests
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/api/v1/ranr/email/", emptyResponce)
	myRouter.HandleFunc("/api/v1/ranr/email/{asn}", GetEmail)
	// asbi requests
	myRouter.HandleFunc("/api/v1/asbi/location/", emptyResponce)
	myRouter.HandleFunc("/api/v1/asbi/location/{locationid}", asbiGetLocation)
	//
	log.Fatal(http.ListenAndServe(":9090", myRouter))
}

// GetEmail - http get email
func GetEmail(w http.ResponseWriter, r *http.Request) {
	// log message
	loggers.Info.Printf("HTTP Request received")
	// determine variables
	vars := mux.Vars(r)
	asn := vars["asn"]
	loggers.Info.Printf("HTTP Request received with AS: %s", asn)
	// get email address by var
	email, err := whois.GetEmailByASN(asn, "w.ranr.noc.gov.ru")
	if err != nil {
		loggers.Error.Printf("Application error:  %v", err)
	}
	// Print email in HTTP Responce
	loggers.Info.Printf("Print email result in HTTP Request: %s", email)
	fmt.Fprintf(w, email)
}

// asbiGetLocation - get location address by location id
func asbiGetLocation(w http.ResponseWriter, r *http.Request) {
	// log message
	loggers.Info.Printf("asbiGetLocation(): HTTP Request received")
	// determine variables
	vars := mux.Vars(r)
	locationID := vars["locationid"]
	loggers.Info.Printf("asbiGetLocation(): HTTP Request received with location id: %s", locationID)
	//
	location, err := asbi.GetLocationByID(locationID)
	//
	if err != nil {
		loggers.Error.Printf("asbiGetLocation(): Error:  %v", err)
		loggers.Info.Printf("Print location: %v", location)
		fmt.Fprintf(w, "")
	} else {
		// convert struct to json
		locationjson, err := json.Marshal(location)
		if err != nil {
			loggers.Error.Printf("asbiGetLocation(): Json creation error:  %v", err)
			fmt.Fprintf(w, "")
			return
		}
		// Print email in HTTP Responce
		loggers.Info.Printf("Print location: %v", location)
		fmt.Fprintf(w, string(locationjson))
	}
}
