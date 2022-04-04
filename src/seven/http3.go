package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type databases map[string]dollors

type dollors float32

func (db databases) list(w http.ResponseWriter, r *http.Request) {
	//for item, price := range db {
	//	fmt.Fprintf(w, "%s: %.2f\n", item, price)
	//}
	templ := template.Must(template.ParseFiles("./src/six/itemSheet.html"))
	if err := templ.Execute(w, db); err != nil {
		log.Println(err)
	}
}

func (db databases) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %s", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%.2f\n", price)
}

func (db databases) update(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	item := query.Get("item")
	price := query.Get("price")
	num, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "illegal price update: %s", price)
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s", item)
		return
	}
	db[item] = dollors(num)
	fmt.Fprintf(w, "success update: %s\n", item)
}

func (db *databases) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	_, ok := (*db)[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s", item)
		return
	}
	delete(*db, item)
	fmt.Fprintf(w, "delete success: %s", item)
}

func (db *databases) create(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	item := query.Get("item")
	price := query.Get("price")
	_, ok := (*db)[item]
	if ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "already have item in database: %s", item)
		return
	}
	priceFloat, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "illeagl price: %s", price)
		return
	}
	(*db)[item] = dollors(priceFloat)
	fmt.Fprintf(w, "create item success: %s", item)
}

func main() {
	db := databases{"shoes": 50, "socks": 5}
	// 请求多路器
	//http.HandleFunc("/list", db.list)
	//http.HandleFunc("/price", db.price)
	mux := http.NewServeMux()
	mux.HandleFunc("/update", (&db).update)
	mux.HandleFunc("/delete", (&db).delete)
	mux.HandleFunc("/create", (&db).create)
	mux.Handle("/list", http.HandlerFunc((&db).list))
	mux.Handle("/item", http.HandlerFunc((&db).price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
