// +build ignore

package inbox

import (
	"fmt"

	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	err = inst.Inbox.Sync()
	e.CheckErr(err)

	i := len(inst.Inbox.Conversations)
	for inst.Inbox.Next() {
		i += len(inst.Inbox.Conversations)
	}
	fmt.Printf("You have %d opened conversations\n", i)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
