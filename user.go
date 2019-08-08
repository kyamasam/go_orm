package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)
//define database connection
var db *gorm.DB
var err error
//define user struct

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}


//create a migration function
//use the orm to create table (sql statements)
func InitialMigration()  {
	//open the db connection
	db,err = gorm.Open("sqlite3","test.db")
	//check the database connection
	if err !=nil{
		fmt.Println(err.Error())
		// The panic built-in function stops normal execution of the current
		// goroutine.
		panic("failed to connect to database")
	}
	//it is idiomatic to defer db.Close() if the sql.DB should not have a lifetime
	// beyond the scope of the function.
	db.Debug().AutoMigrate(&User{})
}
func GetDb() *gorm.DB {
	return db
}
func AllUsers(w http.ResponseWriter , r *http.Request)  {
	var users []User
	//find all users
	GetDb().Find(&users)

	//encode into json
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter , r *http.Request)  {

	vars := mux.Vars(r)
	id := vars["id"]
	var user User
	GetDb().Where("id = ?", id).Find(&user)
	json.NewEncoder(w).Encode(user)

}
//todo create a handler for all kinds of values inside request body :ie both form-data and raw
func NewUser(w http.ResponseWriter , r *http.Request)  {
	decoder :=json.NewDecoder(r.Body)
	//create an instance of the user struct
	var user User
	//decode the request body into a struct and failed if any error occurs
	err := decoder.Decode(&user)
	//check for the error
	if err !=nil{
		fmt.Println(err.Error())
		panic(err)
	}
	GetDb().Create(&User{Name:user.Name, Email:user.Email})
	//fetch the user
	//encode into json
	json.NewEncoder(w).Encode(GetDb().Find(&user))

}

func DeleteUser(w http.ResponseWriter , r *http.Request)  {

	vars := mux.Vars(r)
	id := vars["id"]
	var user User
	GetDb().Where("id = ?", id).Find(&user).Delete(user)
	json.NewEncoder(w).Encode("successfully deleted")

}


func UpdateUser(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var user User

	GetDb().Where("id = ?", id).Find(&user)

	//start getting data passed from post req
	decoder :=json.NewDecoder(r.Body)
	var newUser User
	//decode the request body into a struct and failed if any error occurs
	err := decoder.Decode(&newUser)
	//check for the error
	if err !=nil{
		fmt.Println(err.Error())
		panic(err)
	}
	//replace the passed data with the db data
	user.Name = newUser.Name
	user.Email = newUser.Email
	db.Save(&user)

	// return json of the just edited user
	json.NewEncoder(w).Encode(user)
}
