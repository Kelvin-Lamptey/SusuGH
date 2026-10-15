package main

import "fmt"

var payments map[string]float64
var groups map[string]map[string]float64
var users map[string]string
var group_user map[string]string

var cur_user string


func main(){
	setupApp()
	login()
	for 11 > 2{
		showMenu()
	}
}
func login(){
	fmt.Println("Enter your mobile number to login: ")
	fmt.Scanln(&cur_user)
}
func setupApp(){
	users = make(map[string]string)
	payments = make(map[string]float64)
	groups = make(map[string]map[string]float64)
	group_user = make(map[string]string)
}

func showMenu(){
	message := "You are welcome to the SUSU Group.\nChoose your choice \n 1. Create Group \n 2.Open a Group\n 3.Change account\n\nEnter: "

	var input int
	fmt.Println(message)
	fmt.Scanln(&input)
	if input == 1{
		createGroup()
	}else if input == 2{
		openGroup()
	}else if input == 3{
		login()
	}
}

func createGroup(){
	var groupName string
	var amount float64

	fmt.Println("")
	fmt.Println("Create Group")
	fmt.Println("=============")

	fmt.Print("Enter Group Name: ")
	fmt.Scanln(&groupName)
	fmt.Print("Enter Amount to collect weekly: ")
	fmt.Scanln(&amount)
	groups[groupName] = map[string]float64{ cur_user: amount}
	
	fmt.Println("Group Created!!!")
	// fmt.Println(groups)
}

func openGroup(){
	//add functionality for creating groups
}
	