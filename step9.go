//interface
package main

import (
	"errors"
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

var mutex sync.Mutex

type Poll interface {
	Vote(question string) error
	Add(question string)
	List() []item
}

type item struct {
	question string
	vote     int
}

type poll struct {
	items []item
}

func newPoll() Poll {
	return &poll{make([]item, 0)}
}

func (this *poll) index(question string) int {
	for i := range this.items {
		if this.items[i].question == question {
			return i
		}
	}
	return -1
}

func (this *poll) Vote(question string) error {
	index := this.index(question)
	if index == -1 {
		return errors.New("question does not exist")
	}
	this.items[index].vote++
	return nil
}

func (this *poll) Add(question string) {
	this.items = append(this.items, item{question: question})
}

func (this *poll) List() []item {
	return this.items
}

var thepoll Poll

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	vote := r.FormValue("vote")
	if len(vote) > 0 {
		if err := thepoll.Vote(vote); err != nil {
			fmt.Fprintf(w, "ERROR: %v", err)
		}
	}
	add := r.FormValue("add")
	if len(add) > 0 {
		thepoll.Add(add)
	}
	fmt.Fprintf(w, top)
	fmt.Fprintf(w, tabletop)
	for _, item := range thepoll.List() {
		fmt.Fprintf(w, "<tr><td>%v</td><td>%d</td><td><a href=\"./?vote=%v\">vote</a></td></tr>", item.question, item.vote, item.question)
	}
	fmt.Fprintf(w, tablebottom)
	fmt.Fprintf(w, form)
	fmt.Fprintf(w, bottom)
}

func main() {
	thepoll = newPoll()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
