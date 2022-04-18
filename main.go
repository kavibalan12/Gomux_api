package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type employe struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//slice employe struct
var users []employe

//get single user
func Getone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params id
	//loop through employe and find id
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

//get userall
func Getmany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func CreatUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp employe
	_ = json.NewDecoder(r.Body).Decode(&emp)
	users = append(users, emp)
	json.NewEncoder(w).Encode(emp)

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //find user id
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var emp employe
			_ = json.NewDecoder(r.Body).Decode(&emp)
			users = append(users, emp)
			json.NewEncoder(w).Encode(emp)
			return

		}
		json.NewEncoder(w).Encode(users)
	}

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //find user id
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(users)
	}
}
func main() {

	//init router
	r := mux.NewRouter()

	users = append(users, employe{
		ID:        "1",
		Firstname: "kavi",
		Lastname:  "balan",
	})

	users = append(users, employe{
		ID:        "2",
		Firstname: "mahesh",
		Lastname:  "kumar",
	})
	users = append(users, employe{
		ID:        "3",
		Firstname: "sri",
		Lastname:  "priyan",
	})

	r.HandleFunc("/users/{id}", Getone).Methods("GET")
	r.HandleFunc("/users", Getmany).Methods("GET")
	r.HandleFunc("/users", CreatUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
