package main

import (
	"log"
	"net/http"
)

func main() {
	r := buildRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
