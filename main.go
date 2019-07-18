package main

import (
	"fmt"
	"net/http"
	"github.com/picatz/backdoor"
)

func main() {
	// Need to use the backdoor package at some point in the dependencies,
	// which I'm using directly in the application itself. But, this could be done
	// other ways through a middle-man package that uses the backdoor package
	// without explictly calling it in this application.
	backdoor.FakeFunctionality()

	// Intended application functionality.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	// Keeping it to localhost for testing.
	http.ListenAndServe("127.0.0.1:8080", nil)
}
