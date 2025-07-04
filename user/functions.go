package user

import (
	"fmt"
	"p2p_selling_platform/db"
)

func showMsg(id int) {
	var value, _ = db.AllUsers[id]
	fmt.Println("hello  ", value.Name, "!")
	fmt.Println("1.Buy Products")
	fmt.Println("2.show balance")
	fmt.Println("3.Add Balance")
	fmt.Println("4.Buy by Catagory")
	fmt.Println("5.Post for sell")
	fmt.Println("6.logout")
	fmt.Println("Input your choice")

}

func takeOrder(bal float64, id int) float64 {

	fmt.Println("Input Product Key to buy: ")
	var productId int
	fmt.Scan(&productId)
	value, exist := db.AllItems[productId]

	if exist {
		if value.OwnerId != id {
			if value.Price > bal {
				fmt.Println("Incufficient balance")
				return 0
			} else {
				fmt.Println("Your order done. ")
				return value.Price
			}

		} else {
			fmt.Println("Your own Product")
			return 0
		}
	} else {
		fmt.Println("Invalid ID")
		return 0

		return 0
	}
}
func showCatagory() {
	for key, _ := range db.AllCatagories {
		fmt.Println(key, " ")
	}
}
func catagorySearch() {
	showCatagory()
	var catagory string
	fmt.Println("Input your catagory: ")
	fmt.Scan(&catagory)
	for key, value := range db.AllItems {
		if value.Catagory == catagory {
			fmt.Println("Product key: ", key, "product name : ", value.ProductName, "product price : ", value.Price, " Catagory: ", value.Catagory)
		}
	}
}
func inputProdictid() int {
	var productId int
	fmt.Println("Input your product ID: ")
	fmt.Scan(&productId)
	var _, k = db.AllItems[productId]
	if k {
		fmt.Println("Product Id already exist")
		inputProdictid()
	} else {
		return productId
	}
	return 0
}
func sellPost(id int) {
	showCatagory()
	var Catagory string
	fmt.Println("Input your catagory: ")
	fmt.Scan(&Catagory)
	var _, exist = db.AllCatagories[Catagory]
	if exist {
		ProductId := inputProdictid()
		var Name string
		fmt.Println("Input your product name: ")
		fmt.Scan(&Name)
		var Price float64
		fmt.Println("Input your product price: ")
		fmt.Scan(&Price)
		db.AllItems[ProductId] = db.Product{Name, Catagory, Price, id}
	} else {
		fmt.Println("Unavailable Catagory")
	}

}
