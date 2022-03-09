package gomodule

import "net/http"

func StartWebServer() {
	http.Handle("/", http.FileServer(http.Dir("latein")))
	http.ListenAndServe(":9090", nil)
}
