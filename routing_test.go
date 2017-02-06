package routing

import (
	"bufio"
	"os"
	"testing"
)

func TestRouting__All(t *testing.T) {
	f, err := os.Open("testdata/routing-numbers.txt")
	if err != nil {
		t.Fatalf("error reading testdata/routing-numbers.txt")
	}
	defer f.Close()

	// Read test data line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rtn := scanner.Text()[0:9]
		if len(rtn) > 0 && !Valid(rtn) {
			t.Fatalf("%s not seen as valid routing number", rtn)
		}
	}

	// Fail if the scanner had any errors
	if err = scanner.Err(); err != nil {
		t.Fatalf("%s", err)
	}
}

func TestRouting__Valid(t *testing.T) {
	if !Valid("111000025") {
		t.Fatalf("valid test fail")
	}
}

func TestRouting__Padding(t *testing.T) {
	// Should be 011000015
	if !Valid("11000015") {
		t.Fatal("padding test fail")
	}
}

func TestRouting__Invalid(t *testing.T) {
	if Valid("3211321") {
		t.Fatal("invalid test fail")
	}
}
