package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
//define database connection
var db *gorm.DB
var err error
//define user struct

type User struct {
	gorm.Model
	name string
	email string
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
	defer db.Close()
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter , r *http.Request)  {
	//verify that everything is working
	fmt.Fprintf(w, "all users endpoint hit")
	
}

func NewUser(w http.ResponseWriter , r *http.Request)  {

	fmt.Fprintf(w,"new users endpoint hit")

}

func DeleteUser(w http.ResponseWriter , r *http.Request)  {
	fmt.Fprintf(w,"delete users endpoint hit")

}

func UpdateUser(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w, "Update users endpoint hit")
}
