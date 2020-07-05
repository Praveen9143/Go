package main

import (
	"fmt"
	"sync"
)

// Write a program for Tiny URL
// The program accepts a lengthy url for a website and converts in to a tiny url and stores it
// when you request with tiny url id it will redirct you the the original site representing the tiny URL
// Also our service needs to handle multiple concurrent requests so we need to design to avoid concurrency issues.

type URLholder struct {
 tinyrul map [string]string
 mu  sync.RWMutex
}

func(s *URLholder) Get(keyv string ) string  {

	s.mu.RLock()
	defer s.mu.Unlock()
	url:= s.tinyrul[keyv]

	return url
}

func (s *URLholder) Set(value string) bool{

   s.mu.Lock()
   defer s.mu.Unlock()

   key := getkey(value)
   _,present := s.tinyrul[key]

   if present{

   	return false

   }

   s.tinyrul[key] = value

	return true
}

func getkey (val string) string  {

	key := val[:1]

    return key
}

func main() {
	fmt.Println("Hello..world")
}


