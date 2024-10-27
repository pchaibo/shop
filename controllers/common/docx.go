package common

import (
	"fmt"

	doc "github.com/nguyenthenguyen/docx"
)

func Docx() {
	// Read from docx file
	r, err := doc.ReadDocxFile("./aa.docx")
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)
	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	// Replace like https://golang.org/pkg/strings/#Replace
	docx1.Replace("xxx", "小明99", -1)
	docx1.Replace("ADD", "new_882002", -1)
	//docx1.ReplaceLink("http://example.com/", "https://github.com/nguyenthenguyen/docx", 1)
	//docx1.ReplaceHeader("out with the old", "in with the new")
	docx1.ReplaceFooter("Change This Footer", "new footer")
	docx1.WriteToFile("./ss.docx")
	//time.Sleep(1 * time.Second)
	//post_email() //调用
	fmt.Println("ok")
	r.Close()
}
