// +build ignore

package media

import (
	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	for inst.Inbox.Next() {
		for _, c := range inst.Inbox.Conversations {
			c.Like()
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
