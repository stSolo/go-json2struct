# Go json2struct
Converts raw json data to go structures
## Roadmap
- [x] Read data from variables and save structs to ./structs.go
- [x] Normalize field's names
- [x] Read and write primitive types
- [x] Read and write slices of primitive types
- [x] Read and write structs into the other structs
- [x] Read and write sleces of structs
- [ ] Read and write nested slices of primitive types
- [ ] Read and write nested slices of structs
- [ ] Add flags to read and write from/into custom files
- [ ] Get name of first struct from file
- [ ] Add flags to read and write from/into list of custom files
## Use
Input your data into ./app/main.go in data var:
```go
var data = `example`
```
Run command:
```bash
go run ./app/main.go
```
Output file struct.go will place into root level in the project.

## Emaples
### Example 1
Input data:
```json
{
	"data": [{
	  "type": "articles",
	  "id": 1,
	  "attributes": {
		"title": "JSON:API paints my bikeshed!",
		"body": "The shortest article. Ever.",
		"created": "2015-05-22T14:56:29.000Z",
		"updated": "2015-05-22T14:56:28.000Z"
	  },
	  "relationships": {
		"author": {
		  "data": {"id": 42, "type": "people"}
		}
	  }
	}],
	"included": [
	  {
		"type": "people",
		"id": 42,
		"attributes": {
		  "name": "John",
		  "age": 80,
		  "gender": "male"
		}
	  }
	]
	} 
```
Output:
```go
package main

type Struct2 struct {
	Updated string `json:"updated"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Created string `json:"created"`
}

type Struct5 struct {
	Id   float64 `json:"id"`
	Type string  `json:"type"`
}

type Struct4 struct {
	Data Struct5 `json:"data"`
}

type Struct3 struct {
	Author Struct4 `json:"author"`
}

type Struct1 struct {
	Type          string  `json:"type"`
	Id            float64 `json:"id"`
	Attributes    Struct2 `json:"attributes"`
	Relationships Struct3 `json:"relationships"`
}

type Struct7 struct {
	Name   string  `json:"name"`
	Age    float64 `json:"age"`
	Gender string  `json:"gender"`
}

type Struct6 struct {
	Attributes Struct7 `json:"attributes"`
	Type       string  `json:"type"`
	Id         float64 `json:"id"`
}

type Struct0 struct {
	Data     []Struct1 `json:"data"`
	Included []Struct6 `json:"included"`
}

```
### Example 2
Input:
```json
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
```

Output:
```go
package main

type Struct5 struct {
	Para         string   `json:"para"`
	Glossseealso []string `json:"GlossSeeAlso"`
}

type Struct4 struct {
	Abbrev    string  `json:"Abbrev"`
	Glossdef  Struct5 `json:"GlossDef"`
	Glosssee  string  `json:"GlossSee"`
	Id        string  `json:"ID"`
	Sortas    string  `json:"SortAs"`
	Glossterm string  `json:"GlossTerm"`
	Acronym   string  `json:"Acronym"`
}

type Struct3 struct {
	Glossentry Struct4 `json:"GlossEntry"`
}

type Struct2 struct {
	Title     string  `json:"title"`
	Glosslist Struct3 `json:"GlossList"`
}

type Struct1 struct {
	Glossdiv Struct2 `json:"GlossDiv"`
	Title    string  `json:"title"`
}

type Struct0 struct {
	Glossary Struct1 `json:"glossary"`
}

```
