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
	//add functionality for showing menu
}
func createGroup(){
	//add functionality for creating groups
}
func openGroup(){
	//add functionality for opening a group
}