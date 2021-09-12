package main

import (
	"net/http"
	"time"
)

var WorkQueue = make(chan Work, 100)

func Collector(w http.ResponseWriter, r *http.Request) {

	// we can only be called using post
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// parse duration
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "You must specify a name ", http.StatusBadRequest)
		return
	}

	work := Work{Name: name, Delay: delay}
	WorkQueue <- work
	w.WriteHeader(http.StatusCreated)

}
