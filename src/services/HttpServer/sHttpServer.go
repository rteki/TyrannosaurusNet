package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const HTTP_ADDR = ":8080"

const URL_STRING_ROOT = "/"
const URL_STRING_AUTH = "/auth"
const URL_STRING_RESOURCES = "/resources/"

func (hs *HttpServer) assignHandlers() {
	http.HandleFunc(URL_STRING_ROOT, hs.rootHandler)
	http.HandleFunc(URL_STRING_AUTH, hs.authHandler)
	http.HandleFunc(URL_STRING_RESOURCES, hs.resourcesHandler)
}

func (hs *HttpServer) service() {
	server := &http.Server{Addr: HTTP_ADDR}

	hs.assignHandlers()

	go func() {
		err := server.ListenAndServeTLS("tServer.crt", "tServer.key")
		if err != nil {
			fmt.Println(err)
			hs.Stop()
		}
	}()

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
				fmt.Print("HttpServer: Closing....")
				server.Shutdown(context.TODO())
				return
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (hs *HttpServer) rootHandler(w http.ResponseWriter, r *http.Request) {
	if !hs.isAuth(w, r) {
		http.Redirect(w, r, URL_STRING_AUTH, http.StatusSeeOther)
		return
	}
}
