## 🔹 Beginner Level (warm-up, like even/odd)

### Print Even and Odd Numbers
Two goroutines print even/odd numbers up to N in sequence using channels.

### Ping-Pong
Two goroutines print "Ping" and "Pong" alternately N times.

### Timer-based Worker
Write a goroutine that prints "tick" every second until stopped.

## 🔹 Intermediate Level

### Producer–Consumer Problem
One goroutine produces numbers and pushes into a channel, multiple consumers read and process them.

### Worker Pool
Create a pool of M workers to process N tasks concurrently, wait for all to finish.

### Race Condition Simulation
Simulate a shared bank account with multiple goroutines depositing/withdrawing, first without sync (to see race), then fix with sync.Mutex and sync/atomic.

### Deadlock Demonstration & Avoidance
Write code that causes deadlock with two goroutines and two locks, then fix it using proper lock ordering.

### Pipeline Pattern
Chain multiple goroutines:
e.g., Generate → Square → Filter Even → Print (each stage is a goroutine with channels).

### Cancellation with Context
Start a goroutine doing some work (like downloading chunks), cancel it halfway using context.WithCancel().

## 🔹 Advanced Level

### Rate Limiter
Implement a leaky bucket or token bucket rate limiter using goroutines + time.Ticker.

### Dining Philosophers Problem
Solve the classic concurrency problem with goroutines and sync.Mutex.

### Concurrent Map
Implement a thread-safe map with sync.RWMutex.

### Job Scheduler
Build a cron-like scheduler that executes registered jobs (functions) at given intervals.

### Concurrent Web Crawler
Crawl multiple URLs in parallel with a limit on max concurrent requests.

### Merge Sort with Goroutines
Implement parallel merge sort where subproblems are sorted concurrently.

### Barrier Synchronization
N goroutines must all reach a point before any can proceed (implement using sync.WaitGroup or channels).

### Bounded Buffer Problem
Implement a bounded buffer with channels (simulate blocking when full/empty).