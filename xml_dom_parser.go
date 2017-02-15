package snparse

import (
	"io"
	"io/ioutil"
	"unicode"

	//"golang.org/x/text/encoding/charmap"
	//"golang.org/x/text/transform"
)

func Parse(filename string) (*Document, error) {

	content, readErr := ioutil.ReadFile(filename)
	if readErr != nil {
		return nil, readErr
	}
	// TODO
	// DetectEncoding(content)

	return ParseDocument(string(content))
}

func ParseReader(reader io.Reader, encoding string) {
	if encoding != "UTF-8" {
		//transform.NewReader(reader, enc.NewDecoder())
	}
}

func IsAllowedNameFirstSymbol(c rune) bool {
	return (unicode.IsLetter(c)) || c == '_'
}

func IsAllowedNameSymbol(c rune) bool {
	return (unicode.IsLetter(c)) ||
		unicode.IsDigit(c) ||
		c == '-' ||
		c == '_'
}
