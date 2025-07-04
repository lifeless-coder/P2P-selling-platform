package login

import (
	"fmt"
	"p2p_selling_platform/admin"
	"p2p_selling_platform/user"
)

func takeInput() (int, string) {
	fmt.Print("Enter your Id: ")
	var input int
	fmt.Scan(&input)
	fmt.Print("Enter your Password: ")
	var pass string
	fmt.Scan(&pass)
	return input, pass
}
func LoginMain() {
	fmt.Println("Please login first")
	fmt.Println("1.Login as user")
	fmt.Println("2.Login as Admin")
	fmt.Println("Select your choice")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		id, pass := takeInput()
		exist := userAuth(id, pass)
		if exist {
			fmt.Println("Login success")
			user.UserMain(id)
		} else {
			fmt.Println("Login fail")
		}
	case 2:
		id, pass := takeInput()
		exist := adminAuth(id, pass)
		if exist {
			fmt.Println("Login success")
			admin.AdminMain()
		} else {
			fmt.Println("Login fail")
		}
	default:
		fmt.Println("Invalid choice")
	}
}
