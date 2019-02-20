package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const HTTP_ADDR = ":8080"

func (hs *HttpServer) service() {
	server := &http.Server{Addr: HTTP_ADDR}

	http.Handle("/", hs)

	err := server.ListenAndServeTLS("tServer.crt", "tServer.key")
	if err != nil {
		fmt.Print(err)
		return
	}
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
		time.Sleep(time.Millisecond * 100)
	}
}

func (hs *HttpServer) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	fmt.Printf(req.RemoteAddr)
}
