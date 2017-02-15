package snparse

//	"io/ioutil"
//"strings"
//"encoding"
//	"errors"
//"unicode"

const (
	NODE = iota
	TEXT_NODE
	CDATA
)

type XmlDomParser struct {
}

type INode interface {
	getParent()
	addChild() float64
	addAttribute() float64
}

type Document struct {
	Header   *XmlHeader
	Children []*Node
}

func (doc *Document) addChild(node *Node) {
	if doc.Children == nil {
		doc.Children = make([]*Node, 0)
	}
	doc.Children = append(doc.Children, node)
}

type XmlHeader struct {
	Name     string
	Position Position

	Attributes []Attribute
}

func (header *XmlHeader) addAttribute(attr Attribute) {
	if header.Attributes == nil {
		header.Attributes = make([]Attribute, 0)
	}
	header.Attributes = append(header.Attributes, attr)
}

func (header *XmlHeader) hasAttr(attrName string) bool {
	for i := 0; i < len(header.Attributes); i++ {
		if attrName == header.Attributes[i].Name {
			return true
		}
	}
	return false
}

func (header *XmlHeader) getFirstAttrValue(attrName string) string {
	for i := 0; i < len(header.Attributes); i++ {
		if attrName == header.Attributes[i].Name {
			return header.Attributes[i].Value
		}
	}
	return ""
}

type Position struct {
	Line   int
	Column int
}

type Node struct {
	Name      string
	Namespace string
	Position  Position

	Parent   *Node
	Children []*Node

	Attributes []Attribute

	Type int

	// Only for TextNodes
	Text string

	// GetContents()
}

func (node *Node) addChild(child *Node) {
	if node.Children == nil {
		node.Children = make([]*Node, 0)
	}
	node.Children = append(node.Children, child)
}

func (node *Node) addAttribute(attr Attribute) {
	if node.Attributes == nil {
		node.Attributes = make([]Attribute, 0)
	}
	node.Attributes = append(node.Attributes, attr)
}

type Attribute struct {
	Name      string
	Value     string
	Namespace string
}

func IsLatinCharacter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
