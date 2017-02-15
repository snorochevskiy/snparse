package snparse

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"

	"strings"
)

func getEncoding(encodingName string) encoding.Encoding {

	switch strings.ToLower(encodingName) {
	case "cp1251":
		return charmap.Windows1251
	}
	return nil
}

//https://github.com/djimenez/iconv-go

/*
package main

import (
  "code.google.com/p/go-charset/charset"
  "fmt"
  "io/ioutil"
  "strings"
  "unicode/utf8"
)
import _ "code.google.com/p/go-charset/data" // include the conversion maps in the binary

func main() {
  s := "Hello, \x90\xA2\x8A\x45" // CP932 encoded version of "Hello, 世界"

  r, _ := charset.NewReader("CP932", strings.NewReader(s)) // convert from CP932 to UTF-8
  s2_, _ := ioutil.ReadAll(r)
  s2 := string(s2_)
  fmt.Println(s2)                         // => Hello, 世界
  fmt.Println(len(s2))                    // => 13
  fmt.Println(utf8.RuneCountInString(s2)) // => 9
  fmt.Println(utf8.ValidString(s2))       // => true
}
*/

/*
package main

import (
    "io"
    "os"

    "golang.org/x/text/encoding/charmap"
)

func main() {
    f, err := os.Open("my_isotext.txt")
    if err != nil {
        // handle file open error
    }
    out, err := os.Create("my_utf8.txt")
    if err != nil {
        // handler error
    }

    r := charmap.ISO8859_1.NewDecoder().Reader(f)

    io.Copy(out, r)

    out.Close()
    f.Close()
}
*/

/*
import (
    "code.google.com/p/go.text/transform"
    "code.google.com/p/go.text/encoding/charmap"
)

func main() {

sr := strings.NewReader("Текст в Win1251")
tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
buf, err := ioutil.ReadAll(tr)
if err != err {
 // обработка ошибки
}

s := string(buf) // строка в UTF-8
}
*/
