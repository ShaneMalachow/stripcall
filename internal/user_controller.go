package stripcall

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type UserController struct {
	router *mux.Router
	dep    *DependencyMap
}

var userController UserController

func (con UserController) InitializeController(r *mux.Router, dep *DependencyMap) {
	userController = con
	userController.dep = dep
	userController.router = r
	r.HandleFunc("", GetUsers).Methods("GET")
	r.HandleFunc("", CreateUser).Methods("POST")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := userController.dep.DB
	var users []User
	db.Find(&users)
	res := map[string][]User{
		"users": users,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		HandleError(&w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := userController.dep.DB
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(&w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := &User{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error processing JSON")
		fmt.Println("REQUEST: " + string(requestBody))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.Create(&user)
	_, _ = w.Write([]byte("success"))
}
