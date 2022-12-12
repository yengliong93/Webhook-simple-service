package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	fmt.Println("Start client.")
	reader := bufio.NewReader(os.Stdin)
	
	loop:
	for {
		fmt.Println("Please choose one of the following options: ")
		fmt.Println("[1] Show all employee on the server.")
		fmt.Println("[2] Register an employee")
		fmt.Println("[3] Quit")
		fmt.Print("-> ")

		option, _ := reader.ReadString('\n')
		option = strings.Replace(option, "\n", "", -1)
		switch option {
		case "1":
			show_all_employee()
		case "2":
			register_employee()
		case "3":
			break loop
		default:
			fmt.Println("Invalid option.")
		}
	}
	fmt.Println("Exit client.")
}

func show_all_employee(){
	resp, err := http.Get("http://localhost:5000/show")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
        fmt.Println(err)
    }
	fmt.Println(string(bodyBytes))
}

func register_employee(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter employee Name: ")
	fmt.Print("-> ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	fmt.Println("Enter employee age: ")
	fmt.Print("-> ")
	age, _ := reader.ReadString('\n')
	age = strings.Replace(age, "\n", "", -1)
	fmt.Println("Enter employee gender: ")
	fmt.Print("-> ")
	gender, _ := reader.ReadString('\n')
	gender = strings.Replace(gender, "\n", "", -1)

	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)
	params.Add("gender", gender)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "http://localhost:5000/register", body)
	if err != nil {
		fmt.Print(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
        fmt.Println(err)
    }
	fmt.Println(string(bodyBytes))
}