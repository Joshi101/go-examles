package main

import (
	"fmt"
	"io"
	"net/http"
)

func getFoo(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		fmt.Println("Header", k, "value", v)
		w.Write([]byte("bar"))
	}
}

func postFooBar(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		fmt.Println("Header", k, "value", v)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("body ", data)
	w.Write([]byte("login successful"))

}

func main() {

	http.HandleFunc("/foo", getFoo)
	http.HandleFunc("/foobar", postFooBar)
	http.ListenAndServe("localhost:7292", nil)
}
