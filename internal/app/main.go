package main

import (
	"flag"
	"io"
	"log"
	"mulez/internal/generate"
	"os"
	"text/template"
)

const tmpl = "{{range .}}{{range .}}{{printf \"%-12s\" .}}{{end}}\n{{end}}"

var num = flag.Int("num", 10, "number of exercises")
var file string

func init() {
	flag.StringVar(&file, "out", "", "file name to store result")
}

func main() {
	flag.Parse()
	tbl := generate.Table(*num, 4)
	t := template.Must(template.New("ex").Parse(tmpl))
	var wr io.Writer
	if len(file) > 0 {
		if f, err := os.Create(file); err == nil {
			defer f.Close()
			wr = f
		}
	} else {
		wr = os.Stdout
	}

	if err := t.Execute(wr, tbl); err != nil {
		log.Println(err)
	}
}
