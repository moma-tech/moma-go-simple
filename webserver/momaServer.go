package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"top.moma.go.simple/remote"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func chat(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	charRequest := CharRequest{}
	_ = json.Unmarshal(reqBody, &charRequest)
	fmt.Println(charRequest.Propmt)
	resp := remote.ChatApiCall("token", charRequest.Propmt)
	fmt.Println(resp)
	fmt.Fprintf(w, "%+v", resp)
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/chat", chat)
	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

type CharRequest struct {
	Propmt string `json:"prompt"`
}
