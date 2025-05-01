# Go Load Balancer


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
1. Create multiple simple HTTP servers to test the load balancer. You can use any HTTP server, but here's a simple example using Python's built-in HTTP server:
```bash
# Terminal 1
python3 -m http.server 5001

# Terminal 2
python3 -m http.server 5002

# Terminal 3
python3 -m http.server 5003

# Terminal 4
python3 -m http.server 5004

# Terminal 5
python3 -m http.server 5005
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

### Step 5: Monitor Health Checks
1. The load balancer will automatically perform health checks every 2 seconds (as configured)
2. You can observe the health check logs in the terminal where the load balancer is running
3. If a server becomes unhealthy, it will be automatically removed from the pool
4. When a server becomes healthy again, it will be automatically added back to the pool

## Configuration

// ... existing code ...
