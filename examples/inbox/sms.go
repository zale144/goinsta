// +build ignore

package inbox

import (
	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	err = inst.Inbox.Sync()
	e.CheckErr(err)

	if len(inst.Inbox.Conversations) != 0 {
		err = inst.Inbox.Conversations[0].Send("dfghj")
		e.CheckErr(err)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
