package main

import (
	"fmt"
	"gateway/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/patient", handler.PatientHandler)
	http.HandleFunc("/doctor", handler.DoctorHandler)

	fmt.Println("Gateway server is running on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
