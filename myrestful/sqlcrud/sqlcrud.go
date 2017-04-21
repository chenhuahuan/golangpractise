package sqlcrud

import (
	"fmt"
	"log"

	"database/sql"
	"encoding/json"
)

const goodFeedback string = "true"

type user struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Status   int    `json:"status,omitempty"`
}

//see function name
func GetUserByID(db *sql.DB, userID string) (string, error) {

	//check if nil pointer first!!
	if db == nil {
		return "", fmt.Errorf("null point at args in GetUserByID")
	}

	var name string
	var id string

	st, err := db.Prepare("select id,name from user where id = ?")
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("db.Prepare error : %v", err.Error())
	}

	err = st.QueryRow(userID).Scan(&id, &name)
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("st.QueryRow error: %v", err)
	}

	return name, nil
}

//get Users
func GetUsers(db *sql.DB) (string, error) {

	if db == nil {
		return "", fmt.Errorf("null point at args in GetUser")
	}

	const len = 10
	var users []user
	var result string

	log.Println("Trying to get all users...")
	st, err := db.Prepare("select id,name from user limit ?")
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("db.Prepare error : %v", err.Error())
	}

	rows, errQuery := st.Query(len)
	if errQuery != nil {
		log.Println(errQuery.Error())
		return "", fmt.Errorf("st.Query error : %v", errQuery.Error())
	}

	i := 0
	for rows.Next() {
		var name string
		var id int
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println(err.Error())
			return "", fmt.Errorf("sql Scan func error: %v", err)
		}
		users = append(users, user{ID: id, Name: name})
		b, err := json.Marshal(users[i])
		if err != nil {
			log.Println(err.Error())
			return "", fmt.Errorf("json.Marshal func error: %v", err)
		}
		result += fmt.Sprintf("%s", string(b))
		result += "<br />" //newline in html
		i++
	}

	//A test , to return the fake wrong retuen , to throw back to dispatchHandler
	//	return "",fmt.Errorf("json.Marshal func error: %v",err)
	return result, nil

}

func CreateUserNameID(db *sql.DB, name string, id string) (string, error) {

	if db == nil {
		return "", fmt.Errorf("null point at args in CreateUserNameID")
	}

	log.Println(name, id)
	st, err := db.Prepare("INSERT INTO user(id,name) VALUES(?,?)")
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("db.Prepare error  : %v", err.Error())
	}

	log.Printf("Trying to insert to user table id = %v and name = %v\r\n", id, name)

	var result string
	res, err := st.Exec(id, name)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("st.Exec error : %v", err.Error())
	}
	if res != nil {
		log.Println(res)
		result = goodFeedback
	}

	json, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("json.Marshal error: %v", err)
	}

	log.Println("Done:Update success!")
	return string(json), nil

}

func DeleteUserByID(db *sql.DB, userID string) (string, error) {

	if db == nil {
		return "", fmt.Errorf("null point at args in DelteUserByID")
	}

	var result string

	log.Println(userID)
	st, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		log.Print(err)
		return "", fmt.Errorf("db.Prepare error : %v", err.Error())
	}

	res, err := st.Exec(userID)
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("st.Exec error error: %v", err.Error())
	}
	if res != nil {
		result = goodFeedback
	}

	json, err := json.Marshal(result)
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("json.Marshal error : %v", err.Error())
	}

	return string(json), nil

}

//take a *sql.DB as the arg instead of open it inside the function.
//
func UpdateUserNameByID(db *sql.DB, userID string, userName string) (string, error) {

	if db == nil {
		return "", fmt.Errorf("null point at args in UpdateUserNameByID")
	}

	var result string

	log.Printf("Trying to update user(id = %v) to name %v ...\r\n", userID, userName)
	st, err := db.Prepare("UPDATE user SET name=? WHERE id=?")
	if err != nil {
		return result, fmt.Errorf("db.Prepare error : %v", err)
	}

	res, err := st.Exec(userName, userID)
	if err != nil {
		fmt.Print(err)
	}
	if res != nil { //res != nil ,means the return value is normal.
		result = goodFeedback
		return result, nil
	}

	json, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("json.Marshal error : %v", err)
	}

	return string(json), nil

}
