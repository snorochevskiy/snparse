package snparse

import (
	"strconv"
)

func DetectEncoding(content []byte) string {
	ci := CreateCharIterator(string(content))
	header := parseXmlHeader(ci)

	if header == nil || header.hasAttr("encoding") {
		return ""
	}
	return header.getFirstAttrValue("encoding")
}

func errMsg(line int, column int, msg string) string {
	return "[" + strconv.Itoa(line) + "," + strconv.Itoa(column) + "] " + msg
}
