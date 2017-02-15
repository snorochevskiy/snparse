package snparse

//import (
//	"testing"
//)

//func Test_DetectEncoding_ValidUtf8(t *testing.T) {

//	if testing.Short() {
//		t.Skip("skipping test in short mode.")
//	}

//	enc, err := DetectEncoding([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>"))
//	if err != nil {
//		t.Log(err.Error())
//		t.Fail()
//	}

//	if enc != "UTF-8" {
//		t.Fail()
//	}

//}

//func Test_DetectEncoding_NoEncodingAttribute(t *testing.T) {

//	if testing.Short() {
//		t.Skip("skipping test in short mode.")
//	}

//	enc, err := DetectEncoding([]byte("<?xml version=\"1.0\" ?><some>text</some>"))
//	if err != nil {
//		t.Log(err.Error())
//		t.Fail()
//	}

//	if enc != "" {
//		t.Fail()
//	}

//}
