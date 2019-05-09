package common

import (
	"encoding/json"
	"net/http"
)

// http response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func BuildResponse(code int, msg string, data interface{}) (resp [] byte) {
	var (
		response Response
	)
	response.Code = code
	response.Msg = msg
	response.Data = data
	resp, _ = json.Marshal(response)
	return
}

func BuildDefaultResponse(data interface{}) (resp [] byte) {
	var (
		response Response
	)
	response.Code = Success
	response.Msg = SuccessMsg
	response.Data = data
	resp, _ = json.Marshal(response)
	return
}

func Return(code int, msg []byte, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Write(msg)
}
