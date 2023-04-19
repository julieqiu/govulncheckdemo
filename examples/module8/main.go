// This program takes language tags as command-line
// arguments and parses them.

package main

import (
	"log"

	"github.com/buger/jsonparser"
	"github.com/tidwall/gjson"
)

const data = []byte("{hello: world}")

func main() {
	if !gjson.Valid(data) {
		log.Fatalf("invalid json: %v", data)
	}
	jsonparser.GetString("hello")
}
