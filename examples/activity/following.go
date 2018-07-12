// +build ignore

package activity

import (
	"fmt"

	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	act := inst.Activity.Following()
	e.CheckErr(err)

	for act.Next() {
		fmt.Printf("Stories: %d %d\n", len(act.Stories), act.NextID)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
