package file

import "os"

type File struct {
	File *os.File
}

func (f File) WriteLine(s string) {
	f.File.WriteString(s)
	f.File.WriteString("\n")
}
