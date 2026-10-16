package main

import "fmt"
import "strings"

var payment map[string]float64
var groups map[string]map[int]float64
var users map[string]string
var group_user map[string]string
var payments[]map[string]float64

var cur_pot_user map[string]string
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
	payment = make(map[string]float64)
	groups = make(map[string]map[int]float64)
	group_user = make(map[string]string)
	cur_pot_user = make(map[string]string)
}

func showMenu(){
	message := "\nYou are welcome to the SUSU Group.\nChoose your choice \n 1.Create Group \n 2.Open a Group\n 3.Change account\n\nEnter: "

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
	if groupName == ""{
		fmt.Println("Group name cannot be empty.")
		return
	}
	if _, exists := groups[groupName]; exists{
		fmt.Println("Group already exists.")
		return
	}
	fmt.Print("Enter Amount to collect weekly: ")
	fmt.Scanln(&amount)
	if amount <= 0{
		fmt.Println("Amount must be greater than zero.")
		return
	}
	groups[groupName] = map[int]float64{ 1: amount}

	cur_pot_user[groupName] = cur_user

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
	fmt.Println("Choose your choice.\n 1.Add user to group\n 2.Pay amount\n 3.Close current week")
	fmt.Scanln(&response)
	if response == 1{
		addUserToGroup(name)
	}else if response == 2{
		payAmountForGroup(name)
	}else{
		closeCurrentWeek(name)
	}
}
func addUserToGroup( groupName string){
	if _, exists := groups[groupName]; !exists{
		fmt.Println("Group does not exist.")
		return
	}
	fmt.Println()
	fmt.Println("Enter the user's mobile number to add to the group: ")
	var user string
	fmt.Scanln(&user)
	if user == ""{
		fmt.Println("User cannot be empty.")
		return
	}
	if _, exists := group_user[user]; exists{
		fmt.Println("User is already in a group.")
		return
	}
	group_user[user] = groupName

	fmt.Println(user,"added to", groupName, "group.")

	// fmt.Println(group_user)
}
func payAmountForGroup( groupName string){
	fmt.Println()
	fmt.Println("Enter the amount to pay for the current week")
	var amount float64
	var amount_left float64
	// var active_user string
	var max_amount float64
	fmt.Scanln(&amount)
	for active_cur_week := range groups[groupName]{
		max_amount = groups[groupName][active_cur_week]
		amount_left = max_amount
	}
	if amount <= max_amount && amount > 0{
		for i:=0; i < len(payments);i++{
			for pay:= range payments[i]{
				user := strings.Split(pay,"/|")[0]
				group := strings.Split(pay,"/|")[1]
				amt := payments[i][pay]
				if user == cur_user && group == groupName{
					amount_left = max_amount - amt
				}
			}
		}
		if amount <= amount_left{
			if amount_left == max_amount{
				payments = append(payments, map[string]float64{ cur_user+"/|"+groupName:amount})
			}else{
				for i:=0; i < len(payments);i++{
					for pay:= range payments[i]{
						user := strings.Split(pay,"/|")[0]
						group := strings.Split(pay,"/|")[1]
						if user == cur_user && group == groupName{
							payments[i][pay] += amount
						}
					}
				}
			}
			
		}else{
			fmt.Println("Error with payment")
		}
		
	}else{
		fmt.Println("Error with payment")
	}
}
func getCurrentWeek( groupName string) int{
	 for active_cur_week := range groups[groupName]{
		return active_cur_week
	}
	return 1
}
func closeCurrentWeek( groupName string){
	var active_cur_week int
	active_cur_week = getCurrentWeek(groupName)
	for i := 0; i < len(payments); i++ {
		for pay := range payments[i] {
			delete(payments[i], pay)
		}
	}
	groups[groupName][active_cur_week+1] = groups[groupName][active_cur_week]
	delete(groups[groupName], active_cur_week)


	fmt.Println("Current week closed for group", groupName)
	fmt.Println("Pot paid for", cur_pot_user[groupName], "for group", groupName)
	cur_pot_user[groupName] = getNextPotUser(groupName)

	ReportAllUnfinishedUsersForGroup( groupName)
}

func getNextPotUser( groupName string) string{
	var next_user string
	var use_as_next bool = false
	for user, group := range group_user{
		if use_as_next{
			next_user = user
			break
		}
		if group == groupName && cur_pot_user[groupName] == user{
			use_as_next = true
			continue
		}
	}
	return next_user
}

func ReportAllUnfinishedUsersForGroup( groupName string){
	var active_cur_week int
	active_cur_week = getCurrentWeek(groupName)

	for user, group := range group_user{
		if group == groupName{
			var user_paid float64 = 0
			for i := 0; i < len(payments); i++ {
				for pay := range payments[i] {
					pay_user := strings.Split(pay,"/|")[0]
					pay_group := strings.Split(pay,"/|")[1]
					pay_amount := payments[i][pay]
					if pay_user == user && pay_group == groupName{
						user_paid += pay_amount
					}
				}
			}
			if user_paid < groups[groupName][active_cur_week]{
				fmt.Println("User", user, "owes", groups[groupName][active_cur_week]-user_paid, "for group", groupName)
				fmt.Println("User", user, "has not finished paying for group", groupName)
			}
		}
	}
}