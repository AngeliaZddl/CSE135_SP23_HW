package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
	"os"
)

func main() {
	err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Environment Variables:\n\n")

		fmt.Fprintf(w, "HTTP Request Headers:\n")
		for k, v := range r.Header {
			fmt.Fprintf(w, "%s: %s\n", k, v[0])
		}

		fmt.Fprintf(w, "\nServer Variables:\n")
		for _, e := range os.Environ() {
			fmt.Fprintf(w, "%s\n", e)
		}
	}))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}
