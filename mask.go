package inotify

import "fmt"

type Mask uint32

// The following values are from linux/inotify.h
const (
	// the following are legal, implemented events that user-space can watch for

	IN_ACCESS        Mask = 0x00000001 // File was accessed
	IN_MODIFY        Mask = 0x00000002 // File was modified
	IN_ATTRIB        Mask = 0x00000004 // Metadata changed
	IN_CLOSE_WRITE   Mask = 0x00000008 // Writtable file was closed
	IN_CLOSE_NOWRITE Mask = 0x00000010 // Unwrittable file closed
	IN_OPEN          Mask = 0x00000020 // File was opened
	IN_MOVED_FROM    Mask = 0x00000040 // File was moved from X
	IN_MOVED_TO      Mask = 0x00000080 // File was moved to Y
	IN_CREATE        Mask = 0x00000100 // Subfile was created
	IN_DELETE        Mask = 0x00000200 // Subfile was deleted
	IN_DELETE_SELF   Mask = 0x00000400 // Self was deleted
	IN_MOVE_SELF     Mask = 0x00000800 // Self was moved

	// the following are legal events.  they are sent as needed to any watch
	IN_UNMOUNT    Mask = 0x00002000 // Backing fs was unmounted
	IN_Q_OVERFLOW Mask = 0x00004000 // Event queued overflowed
	IN_IGNORED    Mask = 0x00008000 // File was ignored

	// helper events
	IN_CLOSE Mask = IN_CLOSE_WRITE | IN_CLOSE_NOWRITE // close
	IN_MOVE  Mask = IN_MOVED_FROM | IN_MOVED_TO       // moves

	// special flags
	IN_ONLYDIR     Mask = 0x01000000 // only watch the path if it is a directory
	IN_DONT_FOLLOW Mask = 0x02000000 // don't follow a sym link
	IN_EXCL_UNLINK Mask = 0x04000000 // exclude events on unlinked objects
	IN_MASK_CREATE Mask = 0x10000000 // only create watches
	IN_MASK_ADD    Mask = 0x20000000 // add to the mask of an already existing watch
	IN_ISDIR       Mask = 0x40000000 // event occurred against dir
	IN_ONESHOT     Mask = 0x80000000 // only send event once

	// All of the events - we build the list by hand so that we can add flags in
	// the future and not break backward compatibility.  Apps will get only the
	// events that they originally wanted.  Be sure to add new events here!
	IN_ALL_EVENTS Mask = IN_ACCESS | IN_MODIFY | IN_ATTRIB | IN_CLOSE_WRITE | IN_CLOSE_NOWRITE | IN_OPEN | IN_MOVED_FROM | IN_MOVED_TO | IN_DELETE | IN_CREATE | IN_DELETE_SELF | IN_MOVE_SELF
)

func (wm Mask) String() string {

	var v string

	if wm&IN_ACCESS == IN_ACCESS {
		v += "IN_ACCESS"
	}

	if wm&IN_MODIFY == IN_MODIFY {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_MODIFY"
	}

	if wm&IN_ATTRIB == IN_ATTRIB {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_ATTRIB"
	}

	if wm&IN_CLOSE_WRITE == IN_CLOSE_WRITE {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_CLOSE_WRITE"
	}

	if wm&IN_CLOSE_NOWRITE == IN_CLOSE_NOWRITE {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_CLOSE_NOWRITE"
	}

	if wm&IN_OPEN == IN_OPEN {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_OPEN"
	}

	if wm&IN_MOVED_FROM == IN_MOVED_FROM {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_CLOSE_NOWRITE"
	}

	if wm&IN_MOVED_TO == IN_MOVED_TO {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_MOVED_TO"
	}

	if wm&IN_CREATE == IN_CREATE {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_CREATE"
	}

	if wm&IN_DELETE == IN_DELETE {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_DELETE"
	}

	if wm&IN_DELETE_SELF == IN_DELETE_SELF {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_DELETE_SELF"
	}

	if wm&IN_MOVE_SELF == IN_MOVE_SELF {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_MOVED_TO"
	}

	if wm&IN_ISDIR == IN_ISDIR {
		if len(v) > 0 {
			v += "|"
		}

		v += "IN_ISDIR"
	}

	if len(v) == 0 {
		v = fmt.Sprintf("%d", wm)
	}

	return v
}
