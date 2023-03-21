// This program takes language tags as command-line
// arguments and parses them.

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var update = flag.Bool("update", false, "update .txt and .json files for all modules")

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && strings.HasPrefix(info.Name(), "module") {
			return runGovulncheck(info.Name())
		}
		return nil
	})
}

func runGovulncheck(dir string) error {
	abs, err := filepath.Abs(".")
	if err != nil {
		return err
	}
	abs = abs + "/" + dir
	for _, mode := range []struct {
		filename string
		args     []string
	}{
		{"output.txt", []string{abs}},
		{"verbose.txt", []string{"-v", abs}},
		{"output.json", []string{"-json", abs}},
	} {
		cmd := exec.Command("govulncheck", mode.args...)
		cmd.Dir = abs
		fmt.Printf("Running %q\n", strings.Join(cmd.Args, " "))
		cmd.Stderr = os.Stderr
		out, err := cmd.Output()
		exiterr, _ := err.(*exec.ExitError)
		if err != nil && exiterr.ExitCode() != 3 {
			return fmt.Errorf("govulncheck error: %v", err)
		}
		if err := writeFile(dir+"/"+"output.json", out); err != nil {
			return err
		}
	}
	return nil
}

func writeFile(filename string, output []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(output)
	return err
}
