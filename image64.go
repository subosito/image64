package image64

import (
	"fmt"
	"net/http"
	"encoding/base64"
	"io"
	"io/ioutil"
)

var contentTypes = []string{
	"image/gif",
	"image/png",
	"image/jpeg",
	"image/bmp",
	"image/webp",
	"image/vnd.microsoft.icon",
}

func detectContentType(b []byte) string {
	return http.DetectContentType(b)
}

func isValidContentType(t string) bool {
	for i := range contentTypes {
		if contentTypes[i] == t {
			return true
		}
	}

	return false
}

func encodeToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func format(b []byte) (string, error) {
	ct := detectContentType(b)

	if !isValidContentType(ct) {
		return "", fmt.Errorf("Unsupported content type: %s", ct)
	}

	return fmt.Sprintf("data:%s;base64,%s", ct, encodeToString(b)), nil
}

// Encode converts r into base64 data URI scheme
func Encode(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return format(b)
}

