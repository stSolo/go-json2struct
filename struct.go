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
