package main

import (
	"github.com/andymoon/go-fibonacci"
	"io"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(4) //By changing the number of max procs you can see how the performance changes

	//Takes a pattern and a handler
	http.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		n, err := strconv.ParseInt(r.FormValue("n"), 10, 0)
		if err != nil {
			log.Fatal(err)
		}
		value := fibonacci.Calculate(int(n), false)
		io.WriteString(w, strconv.FormatInt(int64(value), 10))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
