package image64

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var fixtures = []string{
	"test_image.bmp",
	"test_image.gif",
	"test_image.ico",
	"test_image.jpg",
	"test_image.png",
	"test_image.webp",
}

func TestEncode(t *testing.T) {
	for i := range fixtures {
		x := fixtures[i]

		f, err := os.Open(fmt.Sprintf("fixtures/%s", x))
		if err != nil {
			t.Fatal(err)
		}

		e, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s.txt", x))
		if err != nil {
			t.Fatal(err)
		}

		s, err := Encode(f)
		if err != nil {
			t.Fatal(err)
		}

		n := strings.TrimSpace(string(e))

		if n != s {
			t.Fatalf("expected string is not equal: %q != %q", n, s)
		}
	}
}

type badReader struct{}

func (b *badReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("I/O error")
}

func TestEncode_failure(t *testing.T) {
	r := &badReader{}
	_, err := Encode(r)
	if err == nil {
		t.Fatalf("Encode should returns error when r returning error")
	}
}

func TestEncodeFile(t *testing.T) {
	for i := range fixtures {
		x := fixtures[i]

		s, err := EncodeFile(fmt.Sprintf("fixtures/%s", x))
		if err != nil {
			t.Fatal(err.Error())
		}

		e, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s.txt", x))
		if err != nil {
			t.Fatal(err.Error())
		}

		n := strings.TrimSpace(string(e))

		if n != s {
			t.Fatalf("expected string is not equal: %q != %q", n, s)
		}
	}
}

func TestEncodeFile_unsupported(t *testing.T) {
	_, err := EncodeFile("fixtures/test_format.pdf")
	if err == nil {
		t.Fatalf("EncodeFile should returns Unsupported error for application/pdf")
	}
}

func TestEncodeFile_failure(t *testing.T) {
	_, err := EncodeFile("fixture_not_exist")
	if err == nil {
		t.Fatalf("EncodeFile should returns error when file is not exist")
	}
}
