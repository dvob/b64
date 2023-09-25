package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		decode   bool
		url      bool
		padding  bool
		encoding = base64.RawStdEncoding
	)

	flag.BoolVar(&decode, "d", decode, "decode")
	flag.BoolVar(&url, "u", url, "url encoding")
	flag.BoolVar(&padding, "p", padding, "with padding")
	flag.Usage = printUsage
	flag.Parse()

	if url {
		encoding = base64.RawURLEncoding
	}

	if padding {
		encoding.WithPadding(base64.StdPadding)
	} else {
		encoding.WithPadding(base64.NoPadding)
	}

	if decode {
		r := base64.NewDecoder(encoding, os.Stdin)
		_, err := io.Copy(os.Stdout, r)
		return err
	}

	w := base64.NewEncoder(encoding, os.Stdout)
	_, err := io.Copy(w, os.Stdin)
	return err
}

func printUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "b64 reads from standard input prints to standard output.\n\nFlags:\n")
	flag.PrintDefaults()
}
