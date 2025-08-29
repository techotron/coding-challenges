package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Backend server starting...")
	http.HandleFunc("/", func(respWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(respWriter, "Hello from Backend Server")

		reqHostParts := strings.Split(request.Host, ":")

		fmt.Printf("Recieved request from %s\n", request.RemoteAddr)
		fmt.Printf("%s %s %s\n", request.Method, request.RequestURI, request.Proto)
		fmt.Printf("Host: %s\n", reqHostParts[0])
		fmt.Printf("User-Agent: %s\n", request.UserAgent())
		for k, v:= range request.Header {
			fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
		}

		fmt.Println("\nReplied with a hello message")
	})

	http.ListenAndServe("0.0.0.0:8081", nil)
}

