package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// Capture standard output for testing
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestMain(t *testing.T) {
	expectedOutput := "Hello, World!\n"
	actualOutput := captureStdout(main)

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, actualOutput)
	}
}
