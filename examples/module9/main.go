// This program takes language tags as command-line
// arguments and parses them.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v2"
)

type Name struct {
	First string `yaml:"first,omitempty"`
	Last  string `yaml:"last,omitempty"`
}

func main() {
	// 1. Read a JSON or YAML file from the command line.
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("invalid number of args: %v", flag.Args())
	}
	f := flag.Args()[0]
	ext := filepath.Ext(f)
	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	var n Name
	switch ext {
	case ".json":
		// 2a. Parse JSON data.
		if !gjson.Valid(string(data)) {
			log.Fatalf("invalid json: %v", string(data))
		}
		json.Unmarshal(data, &n)
	case ".yaml":
		// 2b. Parse YAML data.
		err := yaml.Unmarshal([]byte(data), &n)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	default:
		log.Fatalf("file format not supported: %q", f)
	}
	fmt.Println(n)
}
