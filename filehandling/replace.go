package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("create.txt")
	if err != nil {
		fmt.Println(err)
	}
	var name, email string
	fmt.Print("Enter name : ")
	fmt.Scan(&name)
	fmt.Print("Enter email : ")
	fmt.Scan(&email)
	mydata := string(data)
	mydata = strings.Replace(mydata, "enter_name", name, -1)
	mydata = strings.Replace(mydata, "enter_email", email, -1)
	err = ioutil.WriteFile("create.txt", []byte(mydata), 0777)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Replaced")
	}
}
