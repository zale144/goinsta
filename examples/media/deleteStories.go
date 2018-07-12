// +build ignore

package media

import (
	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	stories := inst.Account.Stories()

	// you can download item per item or
	// using stories.Delete()
	for stories.Next() {
		for _, item := range stories.Items {
			err = item.Delete()
			e.CheckErr(err)
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
