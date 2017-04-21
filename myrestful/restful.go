package main

import (
	"fmt"
	"log"

	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"demo/myrestful/sqlcrud"
)




//dispatchHandler() function takes the request, and dispatch it to the handle function according to the request type.
func dispatchHandler(response http.ResponseWriter, request *http.Request) {

	var result string

	//sql Open should only once in this function.
	db, err := sql.Open("mysql", "admin:123321@tcp(139.199.152.130:3306)/mydb")
	if err != nil {
		errSQL := fmt.Errorf("SQL Error on sql.Open: %v", err)
		log.Println(errSQL.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error()) // proper error handling instead of panic in your app
		return
	}

	response.Header().Set("Content-type", "application/json")

	err = request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
		return
	}

	params := mux.Vars(request)
	log.Println(request)
	log.Println(params)
	log.Println(request.URL.Path)
	log.Println(params["id"])

	switch request.Method {
	case "GET":
		// GET：https://localhost:12345/users/{id}
		if params["id"] == "" {
			result, err = sqlcrud.GetUsers(db)
			if err != nil {
				log.Println(err.Error())
				http.Error(response, fmt.Sprintf("error getUsers: %v", err), 500)
				return
			}

		} else {
			result, err = sqlcrud.GetUserByID(db, params["id"])
			if err != nil {
				log.Println(err.Error())
				http.Error(response, fmt.Sprintf("error getUserByID %v", err), 500)
				return
			}
		}

	case "POST":
		result, err = sqlcrud.CreateUserNameID(db, request.FormValue("name"), request.FormValue("id"))
		if err != nil {
			log.Println(err.Error())
			http.Error(response, fmt.Sprintf("error createUserNameID: %v", err), 500)
			return
		}

	case "PUT":
		result, err = sqlcrud.UpdateUserNameByID(db, params["id"], request.FormValue("name"))
		if err != nil {
			log.Println(err.Error())
			http.Error(response, fmt.Sprintf("error updateUserNameID: %v", err), 500)
			return
		}

	case "DELETE":
		result, err = sqlcrud.DeleteUserByID(db, params["id"])
		if err != nil {
			log.Println(err.Error())
			http.Error(response, fmt.Sprintf("error deleteUserByID: %v", err), 500)
			return
		}
	default:

	}

	defer db.Close()

	json.NewEncoder(response).Encode(result)

}

func main() {

	/*Usage:

	1）GET:     https://localhost:12345/users
	   Purpose: Get all the info from table user

	2）Get:     https://localhost:12345/{id}
	   Purpose: Get the info from table user =  id

	3）POST:    https://localhost:12345/users/?id=8&name=huahuan8
	   Purpose: Create user id and name

	4）PUT:     https://localhost:12345/users/7?name=huahuan7
	   Purpose: update user whose id =7 to new name

	5）DELETE:     https://localhost:12345/users/8
	   Purpose: Delete user whose id = 8
	*/

	router := mux.NewRouter()
	router.HandleFunc("/users", dispatchHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", dispatchHandler).Methods("GET")
	router.HandleFunc("/users/", dispatchHandler).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", dispatchHandler).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", dispatchHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345", router))

}
