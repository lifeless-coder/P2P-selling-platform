package login

import (
	"p2p_selling_platform/db"
)

//type auth interface {
//	authenticate(name, pass string) bool
//}

func userAuth(id int, pass string) bool {
	user, exist := db.AllUsers[id]
	if exist {
		if user.UserPass == pass {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func adminAuth(id int, pass string) bool {
	user, exist := db.AllAdmins[id]
	if exist {
		if user.AdminPass == pass {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}
