package stripcall

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

type ApiController struct {
	dep *DependencyMap
}

type Controller interface {
	InitializeController(*mux.Router, *DependencyMap)
}

var apiController = &ApiController{}

func ConfigRouter(r *mux.Router, dep *DependencyMap) {
	apiController.dep = dep
	r.HandleFunc("", HelloInfo)
	UserController{}.InitializeController(r.PathPrefix("/users").Subrouter(), dep)
	r.HandleFunc("/events", GetEvents).Methods("GET")
	r.HandleFunc("/calls", GetCalls).Methods("GET")
	r.HandleFunc("/calls", CreateCall).Methods("POST")
	r.HandleFunc("/sms", ReceiveText).Methods("POST")
}

func HelloInfo(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Test 123"))
}

func GetCalls(w http.ResponseWriter, r *http.Request) {
	db := apiController.dep.DB
	var calls []Call
	db.Find(&calls)
	res := map[string][]Call{
		"calls": calls,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func CreateCall(w http.ResponseWriter, r *http.Request) {
	db := apiController.dep.DB
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	call := &Call{}
	err = json.Unmarshal(requestBody, call)
	if err != nil {
		fmt.Println("Error processing JSON")
		fmt.Println("REQUEST: " + string(requestBody))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	call.ReportTime = time.Now()
	db.Create(&call)
	_, _ = w.Write([]byte("success"))
}

func ReceiveText(w http.ResponseWriter, r *http.Request) {
	//db := apiController.dep.DB
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request")
		return
	}
	fmt.Println(string(requestBody))

}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	db := apiController.dep.DB
	var events []Event
	db.Find(&events)
	res := map[string][]Event{
		"events": events,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
