package handler

import "net/http"

func HandleIndex(res http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"get":  "GET",
		"post": "POST",
	}
	switch req.Method {
	case http.MethodGet:
		res.Write([]byte(data["get"]))
	case http.MethodPost:
		res.Write([]byte(data["post"]))
	default:
		http.Error(res, "Bad Request", http.StatusBadRequest)
	}
}
