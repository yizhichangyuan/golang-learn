package main

import (
	"learn/src/first/alias"
	"log"
	"net/http"
	"strconv"
)

func hand(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	//for k, v := range r.Header{
	//	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//}
	//fmt.Fprintf(w, "Host = %q\n", r.Host)
	//fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	//if err := r.ParseForm(); err != nil {
	//	log.Print(err)
	//}
	//for k, v := range r.Form{
	//	fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	//}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, err := strconv.ParseFloat(r.Form["cycles"][0], 64)
	if err != nil {
		log.Print(err)
	}
	alias.Lissajous(w, cycles)
}

func main() {
	http.HandleFunc("/", hand)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
