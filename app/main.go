package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/stSolo/go-json2struct/app/file"
	"github.com/stSolo/go-json2struct/app/writers"
)

var data = `
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}
`

func main() {
	f, err := os.OpenFile("./struct.go", os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fdata := file.File{
		File: f,
	}
	defer fdata.File.Close()
	fdata.WriteLine(`package main`)
	d := make(map[string]interface{})
	temp := []string{}
	json.Unmarshal([]byte(data), &d)

	a := 0
	n := &a
	temp = writers.StringStruct(d, temp, n)
	writers.WriteFile(temp, fdata)

	cmd := exec.Command("gofmt", "-w", "./struct.go")
	cmd.Run()
}
