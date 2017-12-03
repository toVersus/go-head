package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	argv := os.Args
	argc := len(argv)

	if argc != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s Argc: %s\n", argv[0], argv)
		os.Exit(1)
	}

	nlines, err := strconv.Atoi(argv[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	if err := doHead(os.Stdin, nlines); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func doHead(f *os.File, nlines int) error {
	if nlines <= 0 {
		return fmt.Errorf("show nothing about the contents")
	}
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		if nlines > 0 {
			fmt.Printf("%s\n", sc.Text())
		}
		nlines--
	}
	return nil
}
