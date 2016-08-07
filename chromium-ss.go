package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func try(browser string) error {
	var argv0 string

	switch browser {
	case "google-chrome":
		argv0 = "google-chrome"
	default:
		argv0 = "chromium-browser"
	}

	p := filepath.Join(filepath.Join(filepath.Join(os.Getenv("HOME"), ".config"), browser), "SingletonSocket")

	if _, err := os.Stat(p); err != nil {
		return err
	}

	wd, _ := os.Getwd()

	conn, err := net.Dial("unix", p)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = fmt.Fprintf(conn, "START\000%s\000%s\000%s", wd, argv0, strings.Join(os.Args[1:], "\000"))

	return err

}

func main() {
	switch filepath.Base(os.Args[0]) {
	case "google-chrome-ss":
		err := try("google-chrome")
		if err != nil {
			try("chromium")
		}
	default:
		err := try("chromium")
		if err != nil {
			try("google-chrome")
		}
	}

}
