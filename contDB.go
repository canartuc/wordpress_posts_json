package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Write2Json gets filename, serialize struct into beautify JSON and writes to the file
func Write2Json(filename string) {
	allPosts, err := GetAllPosts()
	if err != nil {
		log.Fatal("(ERR) Couldn't get all posts: ", err)
	}

	jsonBeauty, err := json.MarshalIndent(allPosts, "", "   ")
	if err != nil {
		log.Fatal("(ERR) Couldn't beautify the JSON serialization: ", err)
	}

	err = ioutil.WriteFile(filename, jsonBeauty, 0777)
	if err != nil {
		log.Fatalf("(ERR) Couldn't write beautify JSON to the file: %s : %x", filename, err)
	}
}
