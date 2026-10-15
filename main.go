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
	group_user[cur_user] = groupName
	fmt.Println("Group Created!!!")
	// fmt.Println(groups)
}

func openGroup(){
	var name string
	fmt.Println("Enter the Group name: ")
	fmt.Scanln(&name)

	if group_user[cur_user] == name{
		handleGroupSession(name) 
	}else{
		fmt.Println("Your account is not associated with that group or that group does not exist.")
	}
}

func handleGroupSession( name string){
	var response int
	fmt.Println()
	fmt.Println("Welcome to group",name)
	fmt.Println("Choose your choice.\n 1.Add user to group\n 2.Pay amount\n 3.")
	fmt.Scanln(&response)
	if response == 1{
		addUserToGroup(name)
	}else if response == 2{
		payAmountForGroup()
	}
}
func addUserToGroup( groupName string){
	fmt.Println()
	fmt.Println("Enter the user's mobile number to add to the group: ")
	var user string
	fmt.Scanln(&user)
	group_user[user] = groupName

	fmt.Println(user,"added to", groupName, "group.")

	// fmt.Println(group_user)
}
func payAmountForGroup(){
	//add functionality for paying
}