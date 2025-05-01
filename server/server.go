package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1] // get port from command-line
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Served from %s\n", port)
		log.Println(msg)
		w.Write([]byte(msg))
	})

	log.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
