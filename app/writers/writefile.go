package writers

import "github.com/stSolo/go-json2struct/app/file"

func WriteFile(data []string, fdata file.File) {
	if len(data) > 0 {
		for _, v := range data {
			fdata.WriteLine(v)
		}
	}
}
