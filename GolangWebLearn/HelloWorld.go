package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func myHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello vegetable540")

	fmt.Println(r.URL)
	fmt.Println(r.Form)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	icon, err := ioutil.ReadFile("logo.ico")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, icon)
}

func main() {
	http.HandleFunc("/", myHello)
	http.HandleFunc("/favicon.ico", favicon)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}