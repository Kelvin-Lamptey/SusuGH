package main

import "fmt"

var payments map[string]float64
var groups map[string]map[string]float64
var users map[string]string

var cur_user string


func main(){
	setupApp()
	for 11 > 2{
		login()
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
}
func showMenu(){
	message := "You are welcome to the SUSU Group.\nChoose your choice \n 1. Create Group \n 2.Open a Group\nEnter: "

	var input int
	fmt.Println(message)
	fmt.Scanln(&input)
	if input == 1{
		createGroup()
	}else if input == 2{
		openGroup()
	}
}
func createGroup(){
	//add functionality for creating groups
}
func openGroup(){
	//add functionality for opening a group
}