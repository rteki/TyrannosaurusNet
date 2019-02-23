package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (hs *HttpServer) isAuth(w http.ResponseWriter, r *http.Request) bool {
	return false
}

func (hs *HttpServer) authHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("resources/html/auth.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	io.Copy(w, file)
}
