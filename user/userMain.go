package user

import (
	"fmt"
	"p2p_selling_platform/db"
)

func UserMain(id int) {

	val, _ := db.AllUsers[id]
	Bal := val.Balance

	showMsg(id)
	var choice int
	fmt.Scanf("%d", &choice)
	fmt.Println(choice)
	switch choice {
	case 1:
		db.Read()
		var price float64
		price = takeOrder(Bal, id)
		Bal = calculatebal(Bal, price)
		val.Balance = Bal

		UserMain(id)
	case 2:
		fmt.Println("Balance: ", Bal)
		UserMain(id)
	case 3:
		Bal = addBalance(Bal)
		val.Balance = Bal
		UserMain(id)
	case 4:
		catagorySearch()
		var price float64
		price = takeOrder(Bal, id)
		Bal = calculatebal(Bal, price)
		val.Balance = Bal

		UserMain(id)
	case 5:
		sellPost(id)
		UserMain(id)
	case 6:
		return
	default:
		fmt.Println("choice not recognized")
		UserMain(id)
	}
}
