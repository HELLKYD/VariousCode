package gomodule

import (
	"fmt"
	"net/http"
)

func StartTheServer(port string) {
	var finalPort string = ":"
	finalPort += port
	http.HandleFunc("/submit.html", handleSubmit)
	http.Handle("/input.html", http.FileServer(http.Dir("static")))
	http.ListenAndServe(finalPort, nil)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are ready to go")
}
