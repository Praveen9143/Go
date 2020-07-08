// Write a program for Tiny URL
// The program accepts a lengthy url for a website and converts in to a tiny url and stores it
// when you request with tiny url id it will redirect you the the original site representing the tiny URL
// Also our service needs to handle multiple concurrent requests so we need to design to avoid concurrency issues.
package main

import (
	"fmt"
	"net/http"
	"sync"
)

type URLholder struct {
	tinyrul map[string]string
	mu      sync.RWMutex
}

var store = NewUrlholder()
var key string

func NewUrlholder() *URLholder {
	return &URLholder{
		tinyrul: make(map[string]string),
		mu:      sync.RWMutex{},
	}
}

const AddForm = `
	<html><body>
	<form method="POST" action="/add">
    URL: <input type="text" name="url">
    <input type="submit" value="Add">
    </form>
    </html></body>`

func (s *URLholder) Get(keyv string) string {

	s.mu.RLock()
	defer s.mu.RUnlock()
	url := s.tinyrul[keyv]

	return url
}

func (s *URLholder) put(value string) bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	key = getkey(value)
	_, present := s.tinyrul[key]

	if present {

		return false

	}

	s.tinyrul[key] = value

	return true
}

func getkey(val string) string {

	key := val[13:18]

	return key
}

func Add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside add function:")
	url := r.FormValue("url")
	w.Header().Set("Content-Type", "text/html")
	if url == "" {
		fmt.Fprint(w, AddForm)
		return
	}

	keyval := store.put(url)
	var tinyurl string

	if keyval {
		tinyurl = key
	}

	fmt.Fprintf(w, "%s", tinyurl)
}

func Redirect(w http.ResponseWriter, r *http.Request) {

	rkey := r.URL.Path[1:]
	fmt.Println("inside redirect retrieved key", rkey)
	rurl := store.Get(rkey)
	fmt.Println("original url retrieved", rurl)
	if rurl == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, rurl, http.StatusFound)
}

func main() {

	port := ":3000"
	fmt.Println("Starting http server")

	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(port, nil)
	fmt.Println("listening on:", port)
}
