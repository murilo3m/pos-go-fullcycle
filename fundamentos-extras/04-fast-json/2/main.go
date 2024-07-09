package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser

	jsonData := `{"user": {"name": "Mu", "age": 28}}`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	user := value.GetObject("user")
	fmt.Printf("User name: %s\n", user.Get("name"))
	fmt.Printf("User age: %s\n", user.Get("age"))

	userJson := value.Get("user").String()

	var user2 User
	if err := json.Unmarshal([]byte(userJson), &user2); err != nil {
		panic(err)
	}

	fmt.Println(user2.Name, user2.Age)
}
