package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type myMux struct {
}

func (m myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		myHello(w, r)
	}
}

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

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("login/login.html")
		if err != nil {
			fmt.Print(err)
		}

		err = t.Execute(w, nil)
		if err != nil {
			fmt.Print(err)
		}
	} else {
		r.ParseForm()
		fmt.Println(r.Form["usename"])
		fmt.Println(r.Form["pwd"])
	}
}

func main() {
	http.HandleFunc("/", myHello)
	// http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
