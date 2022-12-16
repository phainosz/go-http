package routes

import (
	"encoding/json"
	"fmt"
	"github.com/phainosz/go-http/pkg/utils"
	"net/http"
	"strings"
)

var userUrl = "https://jsonplaceholder.typicode.com/users"

type user struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

// Users faz a integracao com a API jsonplaceholder
func Users(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if id := strings.TrimPrefix(req.URL.Path, "/users/"); id != "" && !strings.Contains(id, "users") {
			getUser(writer, id)
			return
		}
		getUsers(writer)
	default:
		response := utils.Response{Message: "method not allowed"}
		response.JSONResponse(writer, http.StatusMethodNotAllowed)
	}
}

// getUsers retorna uma lista de users
func getUsers(writer http.ResponseWriter) {
	//resp, err := httpClient().Get(userUrl)
	//or
	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	//http.NewRequest allow to manipulate the request (headers, body, params, etc)
	if err != nil {
		response := utils.Response{Message: "Failed to get users."}
		response.JSONResponse(writer, http.StatusBadRequest)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	response, err := utils.HttpClient().Do(req)
	defer response.Body.Close()

	var users []user
	err = json.NewDecoder(response.Body).Decode(&users)
	newUsers, err := json.Marshal(&users)
	if err != nil {
		response := utils.Response{Message: "Failed to unmarshal users."}
		response.JSONResponse(writer, http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(newUsers)
}

// getUser retorna um usuario de acordo com o id
func getUser(writer http.ResponseWriter, id string) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(userUrl+"/%s", id), nil)
	if err != nil {
		response := utils.Response{Message: fmt.Sprintf("Failed to get user %s.", id)}
		response.JSONResponse(writer, http.StatusBadRequest)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	response, err := utils.HttpClient().Do(req)
	defer response.Body.Close()

	var user user
	err = json.NewDecoder(response.Body).Decode(&user)
	newUser, err := json.Marshal(&user)
	if err != nil {
		response := utils.Response{Message: "Failed to unmarshal users."}
		response.JSONResponse(writer, http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(newUser)
}
