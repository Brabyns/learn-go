package main

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
	"os"

)

func main (){
	dbName := "data.db"

	_=os.Remove(dbName)
	

	db, err := sql.Open("sqlite3", dbName)
	if err != nil{
		log.Fatal(err)
	}

	defer func(){
		fmt.Println("closing database connection")
		if err := db.Close(); err != nil{
			log.Printf("Ërror closing database connection: %v", err)
		}
	}()

	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("database connection established")

}