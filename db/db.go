package db

import "fmt"

type Product struct {
	ProductName string
	Catagory    string
	Price       float64
	OwnerId     int
}

type User struct {
	Name     string
	PhnNum   string
	UserId   int
	UserPass string
	Balance  float64
}
type Admin struct {
	Admin     string
	AdminId   int
	AdminPass string
}

var Adu = &User{
	Name:     "Adrita",
	PhnNum:   "01517811919",
	UserId:   1,
	UserPass: "123456",
	Balance:  5000,
}
var Ramim = &User{
	Name:     "Ramim",
	PhnNum:   "015000000",
	UserId:   2,
	UserPass: "123456",
	Balance:  500000,
}

var AllItems = map[int]Product{
	1: Product{
		ProductName: "bed",
		Catagory:    "furniture",
		Price:       35000,
		OwnerId:     1,
	},
	2: Product{
		ProductName: "car",
		Catagory:    "car",
		Price:       3500000,
		OwnerId:     2,
	},
}
var AllUsers = map[int]*User{
	1: Adu,
	2: Ramim,
}
var AllAdmins = map[int]Admin{
	1: Admin{
		Admin:     "Admin",
		AdminPass: "123456",
	},
}
var AllCatagories = map[string]int{
	"furniture": 1,
	"car":       1,
}

func Read() {
	for key, product := range AllItems {
		fmt.Println("Product key: ", key, "product name : ", product.ProductName, "product price : ", product.Price, " Catagory: ", product.Catagory)
	}
}
