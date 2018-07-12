// +build ignore

package account

import (
	"fmt"

	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	fmt.Printf("Profile picture URL: %s\n", inst.Account.ProfilePicURL)

	err = inst.Account.RemoveProfilePic()
	e.CheckErr(err)
	fmt.Printf("After calling func: Profile picture URL: %s\n", inst.Account.ProfilePicURL)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
