package main

import(
	"os"
	"log"
	"path/filepath"
)
func main(){

	// if err := os.Mkdir("Downloads", 0755); err != nil {
	// 	log.Fatal(err) 
	// }

	dir := "Downloads/static/images"
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil{
		log.Fatal(err)
	}

	if err := os.RemoveAll("Downloads"); err != nil {
		log.Fatal(err)
	}
}