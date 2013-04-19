//All types are specified here http://golang.org/ref/spec#Types but here is a short summary
package main

import (
	"fmt"
	"net/http"
)

var top string = `
<html>
	<head>
		<title>My Hackable Question Voter</title>
	</head>
	<body>
`
var (
	tabletop    string = `<table border="1"><tr><td>Question</td><td>Votes</td><td></td></tr>`
	tablebottom string = `</table>`
)

var bottom = `</body></html>`

const form = `
<form name="input" action="." method="get">
	Question: <input type="text" name="add">
	<input type="submit" value="Add">
</form>`

func handler(w http.ResponseWriter, r *http.Request) {
	question := "What is the ultimate question?"
	//var question string = "What is the ultimate question?" //see http://golang.org/ref/spec#Short_variable_declarations
	var votes int = 0
	//votes := 0
	fmt.Fprintf(w, top)
	fmt.Fprintf(w, tabletop)
	fmt.Fprintf(w, "<tr><td>%v</td><td>%d</td><td></td></tr>", question, votes)
	fmt.Fprintf(w, tablebottom)
	fmt.Fprintf(w, form)
	fmt.Fprintf(w, bottom)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

/*

bool

uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

string

array

slice

struct

pointer types

maps

functions

interfaces

channels

*/
