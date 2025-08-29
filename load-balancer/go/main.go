package main

import (
	"fmt"
	"html"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/", func(respWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(respWriter, "Hello, %q", html.EscapeString(request.URL.Path))

		reqHostParts := strings.Split(request.Host, ":")

		fmt.Printf("Recieved request from %s\n", request.RemoteAddr)
		fmt.Printf("%s %s %s\n", request.Method, request.RequestURI, request.Proto)
		fmt.Printf("Host: %s\n", reqHostParts[0])
		fmt.Printf("User-Agent: %s\n", request.UserAgent())
		for k, v:= range request.Header {
			fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
		}
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}

