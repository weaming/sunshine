package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var listen string

func init() {
	listen = *flag.String("l", ":80", "Listen [host]:port, default bind to 0.0.0.0")
	flag.Parse()
}

type Req struct {
	Method string
	Host   string
	URI    string
	Header http.Header
	Body   string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var method string
		if r.Method == "GET" {
			method = green(r.Method)
		} else {
			method = red(r.Method)
		}
		log.Println(method, r.Host, r.RequestURI)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "can't read body", http.StatusInternalServerError)
			return
		}

		req := Req{
			Method: r.Method,
			Host:   r.Host,
			URI:    r.RequestURI,
			Header: r.Header,
			Body:   string(body),
		}
		c, _ := JSON(req)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Server", "https://github.com/weaming/sunshine")
		w.Write(c)
	})
	log.Printf("serve http on %v", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}
