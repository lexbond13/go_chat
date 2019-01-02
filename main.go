package main

import (
	"flag"
	"net/http"
	"log"
	"fmt"
	"lexbond@bitbucket.org/lexbond/chat-go.git"
)

var addr = flag.String("addr", "9090", "http service address")

func showIndex(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w,r, "static/index.html")
}

func main() {
	flag.Parse()

	hub := newHub()
	go hub.Run()

	http.HandleFunc("/", showIndex)
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	serveWs(hub, w, r)
	//})

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Listen serve err: ", err)
	} else {
		fmt.Println("Server running on port: 9090")
	}
}
