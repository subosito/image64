package image64

import (
	"strings"
	"testing"
	"os"
	"io/ioutil"
)

func TestEncode(t *testing.T) {
	f, err := os.Open("fixtures/test_image.png")
	if err != nil {
		t.Fatal(err.Error())
	}

	e, err := ioutil.ReadFile("fixtures/test_image.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	s, err := Encode(f)
	if err != nil {
		t.Fatal(err.Error())
	}

	n := strings.TrimSpace(string(e))

	if n != s {
		t.Fatal("expected string is not equal")
	}
}
