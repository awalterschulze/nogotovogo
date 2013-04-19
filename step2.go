//The handle function will handle to http requests on port 3000 and write a response.
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Writing to the web is like writing to a file.")
}

func main() {
	http.HandleFunc("/", handler)
	//http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Writing to the web is like writing to a file.")
	//})
	http.ListenAndServe(":3000", nil)
}
