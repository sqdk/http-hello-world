package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	serviceName := "undefined"
	if os.Getenv("SERVICE_NAME") != "" {
		serviceName = os.Getenv("SERVICE_NAME")
	}
	
	servicePort := "80"
	if os.Getenv("SERVICE_PORT") != "" {
		servicePort = os.Getenv("SERVICE_PORT")
	}

	r := mux.NewRouter()
	r.HandleFunc("/{hehe:.*}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + serviceName))
		log.Println("Hello", serviceName)
	}).Methods("GET")

	err := http.ListenAndServe(":"+servicePort, r)
	if err != nil {
		log.Println(err)
	}
}
