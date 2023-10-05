package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	init_gateway()
}

func init_gateway() {
	http.HandleFunc("/", HandleRootRequest)

	fmt.Println("Serving the gateway...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
