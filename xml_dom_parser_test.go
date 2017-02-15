package snparse

import (
	"testing"
)

func Test_Parse1Tag1Attr(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	document, err := ParseDocument("<tag1 attr1=\"value1\" />")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(document.Children) != 1 {
		t.Log("Expected 1 Node in document, but got ", len(document.Children))
		t.Fail()
	}

	if document.Children[0].Name != "tag1" {
		t.Log("Expected tag name 'tag', but got " + document.Children[0].Name)
		t.Fail()
	}

	if len(document.Children[0].Attributes) != 1 {
		t.Log("Expected 1 attribute, but got ", len(document.Children[0].Attributes))
		t.Fail()
	}

	attr := document.Children[0].Attributes[0]
	if attr.Name != "attr1" || attr.Value != "value1" {
		t.Log("Expected attribute attr1=value1, but got", attr.Name, "=", attr.Value)
		t.Fail()
	}

}

func Test_Parse2Tags(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	document, err := ParseDocument("<tag1/><tag2/>")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(document.Children) != 2 {
		t.Log("Expected 2 Node on same level in document, but got ", len(document.Children))
		t.Fail()
	}

	document2, err := ParseDocument("<tag1 /><tag2 />")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(document2.Children) != 2 {
		t.Log("Expected 2 Node on same level in document, but got ", len(document2.Children))
		t.Fail()
	}
}

func Test_NestedTag(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	document, err := ParseDocument("<tag1><tag2/></tag1>")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(document.Children) != 1 {
		t.Log("Expected 1 Node on 1st level, but got", len(document.Children))
		t.Fail()
	}
	if len(document.Children[0].Children) != 1 {
		t.Log("Expected 1 nested Node, but got", len(document.Children[0].Children))
		t.Fail()
	}
}

func Test_TextNode(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	document, err := ParseDocument("<tag1>inner text</tag1>")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(document.Children) != 1 {
		t.Log("Expected 1 Node on 1st level, but got", len(document.Children))
		t.Fail()
	}
	if len(document.Children[0].Children) != 1 {
		t.Log("Expected 1 nested Node, but got", len(document.Children[0].Children))
		t.Fail()
	}
	innerNode := document.Children[0].Children[0]
	if innerNode.Type != TEXT_NODE {
		t.Log("Expected nested text node")
		t.Fail()
	}
	if innerNode.Text != "inner text" {
		t.Log("Expected 'inner text' text node, but got: ", innerNode.Text)
		t.Fail()
	}
}

const big_doc = `
<?xml vesion="1.0" encoding="utf8" ?>
< 
parent-1 
	attr1 = "value2" att2
	=
	"value2"
	>
		<_nested_1- />
		< _nested-2 >
			some text
		</_nested-2>
</parent-1>
	< parent2_ />	
`

func Test_BigXml(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	document, err := ParseDocument(big_doc)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(document.Children) != 2 {
		t.Log("Expected 2 Node on same level in document, but got ", len(document.Children))
		t.Fail()
	}
}
