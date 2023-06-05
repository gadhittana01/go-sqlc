package resthttp

import (
	"encoding/json"
	"log"
	"net/http"
)

type baseResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse() baseResp {
	return baseResp{}
}

func (br *baseResp) SetBadRequest(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Bad Request"
	}
	br.Status = "Bad Request"
	br.Message = msg
	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.Data, "setInternalServerError error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(respBytes)
}

func (br *baseResp) SetInternalServerError(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Internal server error"
	}
	br.Status = "Internal Server Error"
	br.Message = msg
	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.Data, "setInternalServerError error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(respBytes)
}

func (br *baseResp) SetOK(data interface{}, w http.ResponseWriter) {
	br.Data = data
	br.Status = "Success"
	br.Message = "Success"

	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.Data, "setOK error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func (br *baseResp) SetCreated(data interface{}, w http.ResponseWriter) {
	br.Data = data
	br.Status = "Success"
	br.Message = "Success"

	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.Data, "setOK error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBytes)
}

func (br *baseResp) SetNotFound(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Not Found"
	}
	br.Status = "Not Found"
	br.Message = msg
	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.Data, "setInternalServerError error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(respBytes)
}
