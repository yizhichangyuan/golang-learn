package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Tracker struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var trackers = Trackers{
	{"Go", "Delilah", "From the Roots Up", 2012, lengths("3m38s")},
	{"Go", "Moby", "Moby", 1992, lengths("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, lengths("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, lengths("4m24s")},
}

func lengths(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type Sign int

const (
	Title Sign = iota
	Artist
	Album
	Year
	Length
)

type Trackers []*Tracker

func (t Trackers) Len() int {
	return len(t)
}

func (t Trackers) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Trackers) Less(i, j int) bool {
	return t[i].Title < t[j].Title
}

type multiOrder struct {
	Trackers
	sign Sign
}

func (m *multiOrder) Init(t Trackers, sign Sign) *multiOrder {
	m.Trackers = t
	m.sign = sign
	return m
}

func (m *multiOrder) SetSign(sign Sign) {
	m.sign = sign
}

// 覆盖其中内部成员Trackers的方法Less
func (m *multiOrder) Less(i, j int) bool {
	switch m.sign {
	case 0:
		return m.Trackers[i].Title < m.Trackers[j].Title
	case 1:
		return m.Trackers[i].Artist < m.Trackers[j].Artist
	case 2:
		return m.Trackers[i].Album < m.Trackers[j].Album
	case 3:
		return m.Trackers[i].Year < m.Trackers[j].Year
	case 4:
		return m.Trackers[i].Length < m.Trackers[j].Length
	default:
		panic("you choose wrong sign")
	}
}

func click(sign Sign) {
	multiorder := new(multiOrder).Init(trackers, sign)
	sort.Sort(multiorder)
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/title":
		click(Title)
	case "/artist":
		click(Artist)
	case "/album":
		click(Album)
	case "/year":
		click(Year)
	case "/length":
		click(Length)
	}
	//dir, _ := os.Getwd()
	// ParFiles use work dir as base dir
	templ := template.Must(template.ParseFiles("./src/six/index.html"))
	if err := templ.Execute(w, trackers); err != nil {
		log.Println(err)
	}
}

func main() {
	//m := new(multiOrder)
	//for _, t := range trackers {
	//	fmt.Println(t.Title)
	//}
	http.HandleFunc("/", handle)
	http.ListenAndServe("localhost:8080", nil)
}
