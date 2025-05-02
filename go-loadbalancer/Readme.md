# Go Load Balancer

A simple yet powerful HTTP load balancer written in Go that distributes incoming traffic across multiple backend servers. This load balancer includes health checking capabilities and implements a round-robin algorithm for request distribution.

## Features

- **Round-Robin Load Balancing**: Distributes requests evenly across available servers
- **Health Checking**: Regularly monitors server health and automatically removes unhealthy servers from the pool
- **Configurable**: Easy configuration through JSON file
- **Concurrent Safe**: Thread-safe implementation using mutexes
- **Reverse Proxy**: Efficiently forwards requests to backend servers

## Prerequisites

- Go 1.16 or higher
- Multiple backend servers to balance load across

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-loadbalancer.git
cd go-loadbalancer
```

2. Build the project:
```bash
go build
```

## Running the Project

### Step 1: Set Up Backend Servers
1. Create multiple Go HTTP servers to test the load balancer. First, create a `server.go` file with the following content:
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port from command line argument or use default
	port := "5001"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Create a simple handler that returns the server's port
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server on port %s\n", port)
	})

	// Start the server
	addr := ":" + port
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
```

2. Build and run multiple instances of the server on different ports:
```bash
# Terminal 1
go build server.go
./server 5001

# Terminal 2
go build server.go
./server 5002

# Terminal 3
go build server.go
./server 5003

# Terminal 4
go build server.go
./server 5004

# Terminal 5
go build server.go
./server 5005
```

### Step 2: Configure the Load Balancer
1. Edit the `config.json` file to match your server setup:
```json
{
  "port": ":8080",
  "healthCheckInterval": "2s",
  "servers": [
    "http://localhost:5001",
    "http://localhost:5002",
    "http://localhost:5003",
    "http://localhost:5004",
    "http://localhost:5005"
  ]
}
```

### Step 3: Run the Load Balancer
1. In a new terminal, start the load balancer:
```bash
./go-loadbalancer
```

### Step 4: Test the Load Balancer
1. Open your web browser and visit `http://localhost:8080`
2. The load balancer will distribute requests across your backend servers
3. You can verify the load balancing by:
   - Checking the `X-Forwarded-Server` header in the response
   - Observing requests being distributed across different server ports
   - Stopping one of the backend servers to see how the load balancer handles failures

4. You can also test using curl commands:
```bash
# Test a single request and check which server handled it
curl -s -i http://localhost:8080 | grep -i "X-Forwarded-Server"

# Make multiple requests to see load balancing in action
for i in {1..5}; do curl -s -i http://localhost:8080 | grep -i "X-Forwarded-Server"; done

# Test with verbose output to see all headers
curl -v http://localhost:8080
```

### Step 5: Monitor Health Checks
1. The load balancer will automatically perform health checks every 2 seconds (as configured)
2. You can observe the health check logs in the terminal where the load balancer is running
3. If a server becomes unhealthy, it will be automatically removed from the pool
4. When a server becomes healthy again, it will be automatically added back to the pool

## Configuration

The load balancer is configured using a `config.json` file with the following structure:

```json
{
  "port": ":8080",
  "healthCheckInterval": "2s",
  "servers": [
    "http://localhost:5001",
    "http://localhost:5002",
    "http://localhost:5003",
    "http://localhost:5004",
    "http://localhost:5005"
  ]
}
```

- `port`: The port on which the load balancer will listen
- `healthCheckInterval`: How often to check server health (e.g., "2s" for 2 seconds)
- `servers`: List of backend server URLs

## Usage

1. Start your backend servers on the configured ports
2. Run the load balancer:
```bash
./go-loadbalancer
```

The load balancer will start listening on the configured port and begin distributing requests to the backend servers.

## How It Works

1. The load balancer reads the configuration from `config.json`
2. It initializes a pool of servers and starts health checking each server
3. When a request comes in:
   - The load balancer selects the next healthy server using round-robin
   - If no healthy servers are available, returns a 503 Service Unavailable response
   - Otherwise, forwards the request to the selected server

## Health Checking

The load balancer performs regular health checks on all servers:
- Sends HEAD requests to each server at the configured interval
- Marks servers as unhealthy if they don't respond with a 200 OK status
- Automatically removes unhealthy servers from the pool
- Re-adds servers to the pool when they become healthy again

## Security Considerations

- The load balancer adds an `X-Forwarded-Server` header to responses, which may expose internal server information
- Consider removing or modifying this header in production environments

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
