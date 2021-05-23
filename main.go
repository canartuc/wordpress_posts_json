package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	posts, err := GetAllPosts()
	if err != nil {
		log.Fatal("(ERR) Sorry, cannot get all the posts: ", err)
	}

	postsJson, err := json.Marshal(posts)
	if err != nil {
		log.Fatal("(ERR) Sorry, cannot convert struct to JSON: ", err)
	}

	fmt.Println(string(postsJson))
}
