package main

import (
	"fmt"
	"net/http"
)

func main() {
	_, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println(err)
	}
}
