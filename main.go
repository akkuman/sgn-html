package main

import "net/http"

func main() {
	http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))
}