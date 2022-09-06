package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

type File struct {
	Path string
}

func (f *File) GetContent() (c []byte, err error) {
	c, err = os.ReadFile(f.Path)
	return
}

func (f *File) GetContentStr() (c string, err error) {
	b, err := f.GetContent()
	if err != nil {
		return
	}
	c = string(b)
	return
}

func (f *File) GetOSFile() (file *os.File, err error) {
	file, err = os.OpenFile(f.Path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		file, err = os.OpenFile(f.Path, os.O_CREATE, os.ModeAppend)
		return
	}
	return
}

func (f *File) WriteAndSave(content map[string]string) (err error) {
	file, err := f.GetOSFile()
	if err != nil {
		return
	}
	b, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Error while converting strings to bytes.", err)
		return
	}
	_, err = file.Write(b)
	if err != nil {
		log.Fatalln("Error while writing to the file.", err)
		return
	}
	return
}

func (f *File) CountFileLines() (count int, err error) {
	str, err := f.GetContentStr()
	if err != nil {
		log.Fatalln("Unable to open the file.", err)
		return
	}
	strArr := strings.Split(str, "")
	for _, v := range strArr {
		if v == "\n" {
			count += 1
		}
	}
	fmt.Printf("Lines count of the file '%v' is '%v'\n", f.Path, count)
	return
}

func ReadJSON(p string) (users Users) {
	f, err := os.Open(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := io.ReadAll(f)
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

func WriteJSON(path, data string) (b []byte) {
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
	f, err := os.Create(path)
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
