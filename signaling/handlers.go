package main

import (
	"encoding/json"
	"net/http"
)

type GetResponse struct {
	uid     string
	host    string
	timeout int64
}

// VersionHttpHandler returns the version of the signaling system
func VersionHttpHandler(w http.ResponseWriter, req *http.Request) {
	logger.Debugf("[recv] %v", req.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(releaseVersion))
}

// GetHttpHandler provides the query function, which allows a client/server
// query the UDP punching IP and port of another client/server
func GetHttpHandler(w http.ResponseWriter, req *http.Request) {
	logger.Debugf("[recv] %v", req.URL.Path)
	token := req.FormValue("token")
	if verifyToken(token) {
		uid := req.FormValue("uid")
		host, timeout := getInfo(uid)
		if host == "" || timeout <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong uid"))
		} else {
			w.WriteHeader(http.StatusOK)
			resp := GetResponse{uid, host, timeout}
			b, err := json.Marshal(resp)
			if err != nil {
				// this should never happen
				panic("something wrong with the use of json")
			}
			w.Write(b)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong token"))
	}
}

// GetHttpHandler allows a client/server set its IP and port for UDP punching
func SetHttpHandler(w http.ResponseWriter, req *http.Request) {
	logger.Debugf("[recv] %v", req.URL.Path)
	token := req.FormValue("token")
	if verifyToken(token) {
		host := req.FormValue("host")
		timeout := req.FormValue("timeout")
		err := setInfo(token, host, timeout)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong token"))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong set parameters"))
	}
}
