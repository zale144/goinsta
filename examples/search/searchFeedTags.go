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

	fmt.Printf("Hello %s!\n", inst.Account.Username)

	// I don't want to make an example of this. Not today.
	tags, err := inst.Search.FeedTags(os.Args[2])
	e.CheckErr(err)

	// TODO

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
