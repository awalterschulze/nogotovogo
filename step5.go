//mutex
package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	top = `
<html>
	<head>
		<title>My Hackable Question Voter</title>
	</head>
	<body>`
	tabletop    = `<table border="1"><tr><td>Question</td><td>Votes</td><td></td></tr>`
	tablebottom = `</table>`
	bottom      = `</body></html>`
	form        = `
<form name="input" action="." method="get">
	Question: <input type="text" name="add">
	<input type="submit" value="Add">
</form>`
)

var questions []string

var mutex sync.Mutex

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	votes := 0
	add := r.FormValue("add")
	if len(add) > 0 {
		questions = append(questions, add)
	}
	fmt.Fprintf(w, top)
	fmt.Fprintf(w, tabletop)
	for _, question := range questions {
		fmt.Fprintf(w, "<tr><td>%v</td><td>%d</td><td></td></tr>", question, votes)
	}
	fmt.Fprintf(w, tablebottom)
	fmt.Fprintf(w, form)
	fmt.Fprintf(w, bottom)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
