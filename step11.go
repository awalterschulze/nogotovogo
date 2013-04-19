//embedding and maps
package main

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
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
	sort.Sort(this)
	return this.items
}

func (this *poll) Len() int {
	return len(this.items)
}

func (this *poll) Swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

func (this *poll) Less(i, j int) bool {
	return this.items[i].vote > this.items[j].vote
}

type LimitVote interface {
	Add(question string)
	VoteWithId(question string, id string) error
	List() []item
	VotesLeft(id string) int
}

type limitVote struct {
	Poll
	voted map[string]int
	limit int
}

func newLimitVote(limit int) LimitVote {
	return &limitVote{newPoll(), make(map[string]int), limit}
}

func (this *limitVote) VoteWithId(question string, id string) error {
	if this.voted[id] >= this.limit {
		return errors.New("You have already voted")
	}
	this.Vote(question)
	this.voted[id]++
	return nil
}

func (this *limitVote) VotesLeft(id string) int {
	return this.limit - this.voted[id]
}

var thepoll LimitVote

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	vote := r.FormValue("vote")
	if len(vote) > 0 {
		if err := thepoll.VoteWithId(vote, r.RemoteAddr); err != nil {
			fmt.Fprintf(w, "ERROR: %v", err)
		}
	}
	add := r.FormValue("add")
	if len(add) > 0 {
		thepoll.Add(add)
	}
	fmt.Fprintf(w, top)
	fmt.Fprintf(w, "VotesLeft = %d", thepoll.VotesLeft(r.RemoteAddr))
	fmt.Fprintf(w, tabletop)
	for _, item := range thepoll.List() {
		fmt.Fprintf(w, "<tr><td>%v</td><td>%d</td><td><a href=\"./?vote=%v\">vote</a></td></tr>", item.question, item.vote, item.question)
	}
	fmt.Fprintf(w, tablebottom)
	fmt.Fprintf(w, form)
	fmt.Fprintf(w, bottom)
}

func main() {
	thepoll = newLimitVote(1000)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
