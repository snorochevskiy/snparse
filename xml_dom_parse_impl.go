package snparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type CharIterator struct {
	Text   string
	Line   int
	Column int

	Runes []rune
	Ind   int
	Size  int
}

func CreateCharIterator(text string) *CharIterator {
	charIterator := new(CharIterator)
	charIterator.Text = text
	charIterator.Line = 1
	charIterator.Column = 1

	charIterator.Runes = []rune(text)
	charIterator.Ind = 0
	charIterator.Size = len(charIterator.Runes)

	return charIterator
}

func (ci *CharIterator) Next() (rune, error) {

	if ci.Ind+1 >= ci.Size {
		return ' ', errors.New("No more symbols left")
	}

	ci.Ind++
	symbol := ci.Runes[ci.Ind]
	if symbol == '\n' {
		ci.Line++
		ci.Column = 1
	} else {
		ci.Column++
	}
	fmt.Println("NEXT: ", string(symbol), "   new Ind:", ci.Ind)
	return symbol, nil
}

func (ci *CharIterator) HasNext() bool {
	return ci.Ind+1 < ci.Size
}

func (ci *CharIterator) Current() rune {
	return ci.Runes[ci.Ind]
}

func (ci *CharIterator) Peek() (rune, error) {
	if ci.Ind+1 >= ci.Size {
		return ' ', errors.New("No more symbols left")
	}
	symbol := ci.Runes[ci.Ind+1]
	return symbol, nil
}

func ParseDocument(xml string) (document *Document, returnErr error) {

	document = new(Document)
	defer func() {
		if msg := recover(); msg != nil {
			document = nil
			returnErr = errors.New(msg.(string))
		}
	}()

	ci := CreateCharIterator(xml)

	skipSpaces(ci)
	document.Header = parseXmlHeader(ci)

	// TODO : add parse DTD

	for {
		skipSpaces(ci)
		if !ci.HasNext() {
			break
		}
		document.addChild(parseXmlNode(ci))
	}

	return document, nil
}

func skipSpaces(ci *CharIterator) {
	if !unicode.IsSpace(ci.Current()) {
		return
	}

	for {
		r, err := ci.Next()
		if err != nil || !unicode.IsSpace(r) {
			return
		}
	}
}

func parseXmlHeader(ci *CharIterator) *XmlHeader {

	if ci.Current() != '<' {

		return nil
	}

	nextRune, err := ci.Peek()
	if err != nil || nextRune != '?' {
		return nil
	}

	header := new(XmlHeader)
	header.Position.Line = ci.Line
	header.Position.Column = ci.Column

	checkedNext(ci) // Move to ?
	checkedNext(ci) // Move to next symbol

	skipSpaces(ci) // If there are spaces before node name

	header.Name = parseName(ci)

	for {
		fmt.Println("---> current: ", string(ci.Current()), "   peek: ", string(checkedPeek(ci)))
		if ci.Current() == '?' {

			if checkedPeek(ci) == '>' {
				checkedNext(ci) // move from ? to >
				checkedNext(ci) // move from > to next symbol
				break
			} else {
				panic(expectedSymbolMsg("parseXmlHeader", '>', ci))
			}
		} else if unicode.IsSpace(ci.Current()) {
			skipSpaces(ci)
		} else if IsAllowedNameFirstSymbol(ci.Current()) {
			header.addAttribute(parseAttribute(ci))
		} else {
			panic(unexpectedSymbolMsg("parseXmlHeader", ci))
		}
	}

	return header
}

func parseNode(ci *CharIterator) *Node {
	skipSpaces(ci)

	// TODO : add cdata
	if ci.Current() == '<' {
		return parseXmlNode(ci)
	} else {
		return ParseTextNode(ci)
	}
}

func parseXmlNode(ci *CharIterator) *Node {
	node := new(Node)
	node.Type = NODE

	skipSpaces(ci)

	if ci.Current() != '<' {
		panic(unexpectedSymbolMsg("parseXmlNode", ci))
	}

	checkedNext(ci)

	skipSpaces(ci)
	node.Name = parseName(ci)

	for {
		if unicode.IsSpace(ci.Current()) {
			skipSpaces(ci)
		}
		if IsAllowedNameFirstSymbol(ci.Current()) {
			node.addAttribute(parseAttribute(ci))
			continue
		}
		break
	}

	hasChildren := true
	if ci.Current() == '/' {
		hasChildren = false
		checkedNext(ci)
		skipSpaces(ci)
	}

	if ci.Current() == '>' {
		if ci.HasNext() {
			checkedNext(ci)
		}
	} else {
		panic(unexpectedSymbolMsg("parseXmlNode", ci))
	}

	if hasChildren {
		for {
			skipSpaces(ci)
			if ci.Current() == '<' && checkedPeek(ci) == '/' {
				// Read all childrens and rached closing tag
				break
			}
			node.addChild(parseNode(ci))
		}
	} else {
		return node
	}

	// Parsing closing tag
	if ci.Current() != '<' {
		panic(expectedSymbolMsg("parseXmlNode", '<', ci))
	}
	if checkedNext(ci) != '/' {
		panic(expectedSymbolMsg("parseXmlNode", '/', ci))
	}
	checkedNext(ci)
	skipSpaces(ci)
	closingName := parseName(ci)
	if closingName != node.Name {
		panic(errMsg2("parseXmlNode", "Unexpected closing tag name "+closingName, ci))
	}
	skipSpaces(ci)
	if ci.Current() != '>' {
		panic(expectedSymbolMsg("parseXmlNode", '>', ci))
	}
	if ci.HasNext() {
		checkedNext(ci)
	}

	return node
}

func ParseTextNode(ci *CharIterator) *Node {
	node := new(Node)
	node.Type = TEXT_NODE

	text := string(ci.Current())
	for {
		symbol, err := ci.Next()
		if err != nil {
			panic(unexpectedEof("ParseTextNode", ci))
		}
		if symbol == '<' {
			break
		}
		text += string(symbol)
	}

	node.Text = strings.TrimSpace(text)

	return node
}

func parseName(ci *CharIterator) string {
	if !IsAllowedNameFirstSymbol(ci.Current()) {
		panic(unexpectedSymbolMsg("readName", ci))
	}
	name := string(ci.Current())

	for {
		c := checkedNext(ci)
		if IsAllowedNameSymbol(c) {
			name += string(c)
		} else {
			break
		}
	}
	return name
}

func parseAttribute(ci *CharIterator) Attribute {
	attrName := parseName(ci)

	skipSpaces(ci)
	if ci.Current() != '=' {
		panic(expectedSymbolMsg("parseAttribute", '=', ci))
	}

	checkedNext(ci)

	skipSpaces(ci)
	if ci.Current() != '"' {
		panic(expectedSymbolMsg("parseAttribute", '"', ci))
	}

	attrValue := ""
	for {
		symbol := checkedNext(ci)
		if symbol == '"' {
			break
		}
		attrValue += string(symbol)
	}

	checkedNext(ci)

	return Attribute{Name: attrName, Value: attrValue}
}

func checkedNext(ci *CharIterator) rune {
	symbol, err := ci.Next()
	if err != nil {
		// TODO : get upper function with reflection?
		panic(unexpectedEof("checkedNext", ci))
	}
	return symbol
}

func checkedPeek(ci *CharIterator) rune {
	symbol, err := ci.Peek()
	if err != nil {
		// TODO : get upper function with reflection?
		panic(unexpectedEof("checkedNext", ci))
	}
	return symbol
}

func unexpectedSymbolMsg(funcName string, ci *CharIterator) string {
	return errMsg2(funcName, "Unexpected symbol "+string(ci.Current()), ci)
}

func expectedSymbolMsg(funcName string, symbol rune, ci *CharIterator) string {
	return errMsg2(funcName, "Eexpected "+string(symbol)+" but got "+string(ci.Current()), ci)
}

func unexpectedEof(funcName string, ci *CharIterator) string {
	return errMsg2(funcName, "Unexpected end of input stream", ci)
}

func errMsg2(funcName string, msg string, ci *CharIterator) string {
	line := strconv.Itoa(ci.Line)
	column := strconv.Itoa(ci.Column)
	return funcName + "[" + line + "," + column + "] " + msg
}
