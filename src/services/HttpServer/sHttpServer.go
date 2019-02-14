package main

import (
	"context"
	"fmt"
	"net/http"
)

func (hs *HttpServer) service() {
	server := &http.Server{Addr: ":8080"}
	// http.HandleFunc("/", httpHandler)
	fs := http.FileServer(http.Dir("./resources/"))
	http.Handle("/", fs)

	go server.ListenAndServe()

	for {
		select {
		case cmd := <-hs.ctrl:
			{
				hs.ctrls[cmd.Id](cmd.Params)
			}
		default:
			{
			}
		case <-hs.stopchan:
			{
				fmt.Print("Closing....")
				server.Shutdown(context.TODO())
				return
			}
		}
	}
}

// func httpHandler(response http.ResponseWriter, request *http.Request) {
// 	file, err := os.Open("../resources/html/home.html")
// 	request.
// 	if err != nil {
// 		fmt.Printf("Cant open html...\n")
// 	}

// 	io.Copy(response, file)

// 	file.Close()
// }
