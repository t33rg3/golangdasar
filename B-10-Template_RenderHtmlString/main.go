package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var view string = `<html>
<head>
	<title>Template</title>
</head>
<body>
	<h1>Hello</h1>
</body>
</html>`

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
	})

	fmt.Println("server start at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
