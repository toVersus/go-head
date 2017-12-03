package main

import (
	"os"
	"testing"
)

func TestDoHeadWithInvalidInput(t *testing.T) {
	path := ".\\content.txt"
	nlines := -1

	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	err = doHead(f, nlines)
	if err.Error() != "show nothing about the contents" {
		t.Error(err)
	}

	defer f.Close()
}

func TestDoHead(t *testing.T) {
	path := ".\\content.txt"
	nlines := 3

	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	err = doHead(f, nlines)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
}
