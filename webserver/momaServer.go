package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"top.moma.go.simple/mlogger"
	"top.moma.go.simple/remote"
)

const ENV_TOKEN = "MOMA_TOKEN"

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func chat(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	charRequest := CharRequest{}
	_ = json.Unmarshal(reqBody, &charRequest)
	mlogger.InfoLogger.Println(charRequest.Propmt)
	mlogger.InfoLogger.Println(os.Getenv(ENV_TOKEN))
	resp := remote.ChatApiCall(os.Getenv(ENV_TOKEN), charRequest.Propmt)
	mlogger.InfoLogger.Println(resp)
	fmt.Fprintf(w, "%+v", resp)
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/chat", chat)
	mlogger.ErrorLogger.Fatal(http.ListenAndServe(":9999", myRouter))
	mlogger.InfoLogger.Println("Server Start at 9999")

}

type CharRequest struct {
	Propmt string `json:"prompt"`
}
