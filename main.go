package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func router(u string) {

	fmt.Println(u)

}

func main() {

	var port int16 = 8080
	var ip = "127.0.0.1"
	var addr = fmt.Sprintf("%s:%d", ip, port)

	fmt.Println("Running on:", addr)

	var reqHandler = func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, req.URL.Path)
		router(req.URL.Path)
	}

	http.HandleFunc("/", reqHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
