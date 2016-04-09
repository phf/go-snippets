// Blah is a silly net/http server that sometimes helps me debug things.
package main

import (
	"flag"
	"github.com/kr/pretty"
	"html/template"
	"log"
	"net/http"
	"time"
)

const html = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>{{ .title }}</title>
	</head>
	<body>
		<h1>{{ .title }}</h1>
		<p>{{ .now }}</p>
		<h2>http.Request</h2>
		<pre>{{ .request }}</pre>
		<h2>http.ResponseWriter</h2>
		<pre>{{ .writer }}</pre>
	</body>
</html>`

var t = template.Must(template.New("blah").Parse(html))
var port = flag.String("p", "8080", "port to listen on")
var verbose = flag.Bool("v", false, "print to stdout as well")

func handler(w http.ResponseWriter, r *http.Request) {
	rs := pretty.Sprintf("%# v\n", r)
	ws := pretty.Sprintf("%# v\n", w)

	if *verbose {
		log.Println(rs)
		log.Println(ws)
	}

	data := map[string]interface{}{
		"now":     time.Now(),
		"request": rs,
		"title":   "net/http debug server",
		"writer":  ws,
	}

	err := t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
