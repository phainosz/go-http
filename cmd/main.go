package main

import (
	"fmt"
	"github.com/phainosz/go-http/pkg/routes"
	"github.com/phainosz/go-http/pkg/utils"
	"log"
	"net/http"
	"strings"
)

var (
	port = ":8000"
)

func main() {
	log.Println("Start application")

	http.HandleFunc("/home/", home)
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", routes.Users)
	http.HandleFunc("/users/", routes.Users)

	log.Fatal(http.ListenAndServe(port, nil))
}

func home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		if id := strings.TrimPrefix(req.URL.Path, "/home/"); id != "" && !strings.Contains(id, "home") {
			response := utils.Response{Message: fmt.Sprintf("id: %s", id)}
			response.JSONResponse(res, http.StatusOK)
			return
		}
		response := utils.Response{Message: "ok"}
		response.JSONResponse(res, http.StatusOK)
	case "POST":
		response := utils.Response{Message: "created"}
		response.JSONResponse(res, http.StatusCreated)
	case "PUT":
		response := utils.Response{Message: "updated"}
		response.JSONResponse(res, http.StatusOK)
	case "DELETE":
		response := utils.Response{Message: "deleted"}
		response.JSONResponse(res, http.StatusOK)
	default:
		response := utils.Response{Message: "method not allowed"}
		response.JSONResponse(res, http.StatusMethodNotAllowed)
	}
}
