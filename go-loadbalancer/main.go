package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"time"
)

// Current field to keep track of which server should handle next request
// The Mutex ensures that our code is safe to use concurrently.
type LoadBalancer struct {
	Current int
	Mutex   sync.Mutex
}

// Server struct to hold multiple server instances
// URL for each server
// isHealthy flag indicates whether the server is available to handle requests
type Server struct {
	URL       *url.URL
	isHealthy bool
	Mutex     sync.Mutex
}

// Configuring the Load Balancer
type Config struct {
	Port                string   `json:"port"`
	HealthCheckInterval string   `json:"healthcheckinterval"`
	Servers             []string `json:"servers"`
}

func loadConfig(file string) (Config, error) {
	var config Config

	data, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err.Error())
	}

	healthCheckInterval, err := time.ParseDuration(config.HealthCheckInterval)
	if err != nil {
		log.Fatalf("Invalid health check interval: %s", err.Error())
	}

	var servers []*Server
	for _, serverUrl := range config.Servers {
		u, _ := url.Parse(serverUrl)
		server := &Server{URL: u, isHealthy: true}
		servers = append(servers, server)
		go healthCheck(server, healthCheckInterval)
	}

	lb := &LoadBalancer{Current: 0}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		checkServer(w, r, lb, servers)
	})

	log.Println("Starting load balancer on port", config.Port)
	err = http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatalf("Error starting load balancer: %s\n", err.Error())
	}
}

func checkServer(w http.ResponseWriter, r *http.Request, lb *LoadBalancer, servers []*Server) {

	server := lb.getNextServer(servers)
	if server == nil {
		http.Error(w, "No healthy server available", http.StatusServiceUnavailable)
		return
	}

	// adding this header just for checking from which server the request is being handled.
	// this is not recommended from security perspective as we don't want to let the client know which server is handling the request.
	w.Header().Add("X-Forwarded-Server", server.URL.String())
	server.ReverseProxy().ServeHTTP(w, r)
}

func (lb *LoadBalancer) getNextServer(servers []*Server) *Server {

	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	for i := 0; i < len(servers); i++ {
		idx := lb.Current % len(servers)
		nextServer := servers[idx]
		lb.Current++

		nextServer.Mutex.Lock()
		isHealthy := nextServer.isHealthy
		nextServer.Mutex.Unlock()

		if isHealthy {
			return nextServer
		}
	}

	return nil
}

func healthCheck(s *Server, hcIntv time.Duration) {

	for range time.Tick(hcIntv) {

		res, err := http.Head(s.URL.String())

		s.Mutex.Lock()

		if err != nil || res.StatusCode != http.StatusOK {
			fmt.Printf("%s is down\n", s.URL)
			s.isHealthy = false
		} else {
			s.isHealthy = true
		}

		s.Mutex.Unlock()
	}
}

func (s *Server) ReverseProxy() *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(s.URL)
}
