package controllers

import "net/http"

func GetControllers() map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"/": func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("test")) },
	}
}
