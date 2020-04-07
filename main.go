package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index page")

}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":9090", nil)

}
