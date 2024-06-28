package inotify

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/g0rbe/go-bytebuilder"
	"golang.org/x/sys/unix"
)

// Flags for Init()
const (
	// Set the O_NONBLOCK file status flag on the open file
	// description (see open(2)) referred to by the new file
	// descriptor.  Using this flag saves extra calls to fcntl(2)
	// to achieve the same result.
	IN_NONBLOCK int = 0x800

	// Set the close-on-exec (FD_CLOEXEC) flag on the new file
	// descriptor.  See the description of the O_CLOEXEC flag in
	// open(2) for reasons why this may be useful.
	IN_CLOEXEC int = 0x80000
)

type Watcher struct {
	fd  int
	wds map[int]string
}

func Init(flags int) (*Watcher, error) {

	fd, err := unix.InotifyInit1(flags)
	if err != nil {
		return nil, err
	}

	inf := new(Watcher)

	inf.fd = fd
	inf.wds = make(map[int]string)

	return inf, nil
}

func (w *Watcher) AddWatch(path string, mask Mask) (int, error) {

	wd, err := unix.InotifyAddWatch(w.fd, path, uint32(mask))
	if err != nil {
		return 0, err
	}

	w.wds[wd] = path

	return wd, nil
}

func (w *Watcher) RmWatch(wd int) error {

	_, err := unix.InotifyRmWatch(w.fd, uint32(wd))

	return err
}

func (w *Watcher) Close() error {

	return unix.Close(w.fd)
}

func (i *Watcher) Read() ([]Event, error) {

	buf := make([]byte, unix.SizeofInotifyEvent+4096)

	n, err := unix.Read(i.fd, buf)
	if err != nil {
		return nil, err
	}

	if n < unix.SizeofInotifyEvent {
		return nil, fmt.Errorf("invalid length: %d", n)
	}

	buf = buf[:n]

	var v []Event

	for len(buf) > 0 {

		wd, ok := bytebuilder.ReadInt32(&buf)
		if !ok {
			return nil, fmt.Errorf("failed to read wd")
		}

		mask, ok := bytebuilder.ReadUint32(&buf)
		if !ok {
			return nil, fmt.Errorf("failed to read mask")
		}

		cookie, ok := bytebuilder.ReadUint32(&buf)
		if !ok {
			return nil, fmt.Errorf("failed to read cookie")
		}

		nlen, ok := bytebuilder.ReadUint32(&buf)
		if !ok {
			return nil, fmt.Errorf("failed to read len")
		}

		var name []byte
		if nlen > 0 {
			name = bytebuilder.ReadBytes(&buf, int(nlen))
			if name == nil {
				return nil, fmt.Errorf("failed to read name")
			}

			name = bytes.TrimSuffix(name, []byte{0x0})
		}

		e := Event{Mask: Mask(mask), Cookie: int(cookie), Name: filepath.Join(i.wds[int(wd)], string(name))}

		v = append(v, e)
	}

	return v, nil
}
