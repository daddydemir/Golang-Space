package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func stat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func ping(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Europe/Istanbul")
	u, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	params := u.Query()
	unix := params.Get("time")
	resp := &response{
		Endtime: time.Duration(time.Now().In(loc).UnixMilli()),
	}
	i64, _ := strconv.ParseInt(unix, 10, 64)
	log.Println(time.Now().UnixMilli()-i64, " ms")
	json.NewEncoder(w).Encode(&resp)
}

type response struct {
	Endtime time.Duration
}
