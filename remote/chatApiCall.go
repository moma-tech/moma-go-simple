package remote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const URL = "https://api.openai.com/v1/chat/completions"
const METHOD = "POST"
const HEADER_AUTH = "Authorization"
const HEADER_CONTENT_TYPE = "Content-Type"
const HEADER_CONTENT_TYPE_VALUE = "application/json"
const MODEL = "gpt-3.5-turbo"

func ChatApiCall(token, prompt string) (output string) {
	output = "null"
	requestBody := RequestBody{}
	requestBody.createBody(prompt)
	jsonData, _ := json.Marshal(&requestBody)
	fmt.Printf("json data: %s\n", jsonData)
	client := &http.Client{}
	req, _ := http.NewRequest(METHOD, URL, bytes.NewBuffer(jsonData))
	req.Header.Set(HEADER_AUTH, token)
	req.Header.Set(HEADER_CONTENT_TYPE, HEADER_CONTENT_TYPE_VALUE)
	response, _ := client.Do(req)
	body, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		errorResponse := ErrorResponse{}
		_ = json.Unmarshal(body, &errorResponse)
		fmt.Printf(errorResponse.Error.Message)
		output = string(errorResponse.Error.Message)
	} else {
		fmt.Print(string(body))
		output = string(body)
	}
	return output
}

func (obj *RequestBody) createBody(prompt string) {
	fmt.Println(prompt)
	obj.Model = MODEL
	var message = Message{Role: "user", Content: prompt}
	obj.Messages = append(obj.Messages, message)
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param"`
		Code    string `json:"code"`
	} `json:"error"`
}
