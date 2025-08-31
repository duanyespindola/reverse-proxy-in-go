    
# Go Reverse Proxy Load Balancer Learning Project

I'm building this project to learn:

1.  **HTTP Fundamentals:** Try to understand HTTP requests o, responses, headers, and status codes on Go with the `net/http` package
2.  **Concurrency:** Deal with multiple incoming client requests and outgoing requests to backend servers simultaneously, and figure out the built-in concurrency model eith  goroutines and channels.
5.  **Project Structure & Modularity:** Test a project separation into logical components (e.g., server, proxy, health checker, balancer strategy).

## Project Scope

This project aims to build a basic reverse proxy load balancer in Go, demonstrating core concepts of HTTP handling, concurrency, and load balancing strategies.

The project will consist of two main components:

1.  **Backend Servers (Echo Servers):**
    *   Simple HTTP servers that listen on a specified port (e.g., `8081`, `8082`, `8083`).
    *   Each server will be identifiable by a number passed as a command-line argument (e.g., `--number=1`).
    *   When hit, a backend server will respond with a unique message, such as "Hello world from server N", where 'N' is its assigned number.
    *   These servers will be launched independently (e.g., `go run server.go --number=1 --port=8081`).

2.  **Reverse Proxy Load Balancer:**
    *   A single Go application that will always listen on a fixed port (e.g., `9090`).
    *   It will be configured with a list of backend server URLs.
    *   When a client makes a request to the load balancer (e.g., `localhost:9090`), the load balancer will forward that request to one of the available backend servers.
    *   The load balancing strategy implemented will be **Round Robin**. This means each successive client request will be routed to the next available backend server in a cyclical order.
    *   The client should receive the response from the chosen backend server (e.g., "Hello world from server 1", then "Hello world from server 2", and so on, cyclically).

  