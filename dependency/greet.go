package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet greets via name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func myGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Zahid")
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(myGreetHandler))
}
