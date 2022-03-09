package gomodule

import "net/http"

func StartWebService() {
	http.Handle("/", http.FileServer(http.Dir("raspi")))
	http.ListenAndServe(":9090", nil)
}
