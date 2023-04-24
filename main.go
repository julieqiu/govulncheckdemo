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
	return filepath.Walk("examples", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), "module") {
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
	abs = abs + "/examples/" + dir
	for _, mode := range []struct {
		filename string
		args     []string
	}{
		{"output.txt", []string{"./..."}},
		{"verbose.txt", []string{"-v", "./..."}},
		{"output.json", []string{"-json", "./..."}},
	} {
		cmd := exec.Command("govulncheck", mode.args...)
		cmd.Dir = abs
		fmt.Printf("%q: Running %q\n", cmd.Dir, strings.Join(cmd.Args, " "))
		cmd.Stderr = os.Stderr
		out, err := cmd.Output()
		if exiterr, ok := err.(*exec.ExitError); ok {
			if err != nil && exiterr.ExitCode() != 3 {
				return fmt.Errorf("govulncheck error: %v", err)
			}
		} else if err != nil {
			return err
		}
		if err := writeFile(abs+"/"+mode.filename, out); err != nil {
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
