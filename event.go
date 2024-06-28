package inotify

import "fmt"

type Event struct {
	Mask   Mask
	Cookie int
	Name   string
}

func (e Event) String() string {

	if e.Cookie == 0 {
		return fmt.Sprintf("%s %s", e.Name, e.Mask)
	}

	return fmt.Sprintf("%s %s %d", e.Name, e.Mask, e.Cookie)

}
