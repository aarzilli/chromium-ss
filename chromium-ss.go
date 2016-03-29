package main

import (
	"net"
	"os"
	"path/filepath"
	"fmt"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	p := os.Getenv("HOME")
	var browser string
	var argv0 string

	switch filepath.Base(os.Args[0]) {
	case "google-chrome-ss":
		browser = "google-chrome"
		argv0 = "google-chrome"
	default:
		browser = "chromium"
		argv0 = "chromium-browser"
	}

	for _, c := range []string{ ".config", browser, "SingletonSocket" } {
		p = filepath.Join(p, c)
	}

	wd, _ := os.Getwd()

	conn, err := net.Dial("unix", p)
	must(err)
	defer conn.Close()

	fmt.Fprintf(conn, "START\000%s\000%s\000%s", wd, argv0, strings.Join(os.Args[1:], "\000"))
}
