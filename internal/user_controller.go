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
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := userController.dep.DB
	var users []User
	db.Find(&users)
	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := userController.dep.DB
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
