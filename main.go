package main

import (
	"fmt"
	"html/template"
	"log"
	"mulez/internal/generate"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const tmpl = `<table style='width: 100%; font-size: x-large;'>
{{range .}}
<tr>{{range .}}<td style='padding: 6px'><b>{{.}}</b></td>{{end}}</tr>
{{end}}
</table>`

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
		log.Println("No PORT environment variable detected, defaulting to", port)
	}
	return ":" + port
}

func handler(w http.ResponseWriter, r *http.Request) {
	num := extractParameters(r.URL.Query())
	txt := template.Must(template.New("exercises").Parse(tmpl))
	t := generate.Table(num, 4)
	if err := txt.Execute(w, t); err != nil {
		fmt.Println(err)
	}
}

func extractParameters(values url.Values) int {
	const defaultValue = 10
	s := values.Get("num")
	if len(s) == 0 {
		log.Println("set number of exercises to the default value:", defaultValue)
		return defaultValue
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		log.Println("set number of exercises to the default value:", defaultValue)
		return defaultValue
	}

	return num
}
