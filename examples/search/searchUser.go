// +build ignore

package search

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("<query>")
	e.CheckErr(err)

	res, err := inst.Search.User(os.Args[2])
	e.CheckErr(err)

	for _, user := range res.Users {
		fmt.Printf("    %s\n", user.Username)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
