package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
