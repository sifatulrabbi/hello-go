package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int64  `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func ReadJSON(p string) (users Users) {
	f, err := os.Open(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	if err := json.Unmarshal(b, &users); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users.Users)
	return
}

func WriteJSON(p string) (b []byte) {
	d := User{
		Name: "Sifatul Rabbi",
		Age:  21,
		Type: "Author",
		Social: Social{
			Facebook: "https://facebook.com/",
			Twitter:  "https://twitter.com/",
		},
	}
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	f, err := os.Create(p)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	_, err = f.Write(b)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return
}

func CopyJSON(from, to string) (done bool) {
	c := ReadJSON(from)
	done = false
	f, err := os.Create(to)
	if err != nil {
		fmt.Printf("Unable to create file, %v\n", err)
		return
	}
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Unable to read file, %v\n", err)
		return
	}
	f.Write(b)
	return true
}
