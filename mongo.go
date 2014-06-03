package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	db *mgo.Session
)

type User struct{
	Name string
	Email string
	Description string
}

func PanicIf(err error){
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func setupDB() *mgo.Session {
	db, err := mgo.Dial("localhost")
	PanicIf(err)
	return db
}

func insertData() {
	coll := db.DB("Pravin").C("User")
	fmt.Println("Collection:", coll)

	err := coll.Insert(&User{"Pravin", "pravinmishra88@gmail.com", "Hey, how are you?"})
	PanicIf(err)
}

func getData() {
	coll := db.DB("Pravin").C("User")
	fmt.Println("Collection data:", coll)

	users := []User{}
	err := coll.Find(bson.M{}).All(&users)
	PanicIf(err)

	fmt.Println("All users: ",users)
}

func main() {
	db = setupDB()
	fmt.Println("Db: ", db)

	insertData()
	getData()
}
