package flock

import (
	"os"
	"sync"

	"github.com/gofrs/flock"
	"golang.org/x/sys/unix"
)

type Flock struct {
	f      *flock.Flock
	fh     *os.File
	locked bool
	m      sync.RWMutex
}

func New(path string) *Flock {
	if Reference {
		return &Flock{f: flock.New(path)}
	}
	var fh, _ = os.Open(path)
	return &Flock{fh: fh}
}
func (f *Flock) Lock() error {
	if Reference {
		return f.f.Lock()
	}
	f.m.Lock()
	defer f.m.Unlock()
	if f.locked {
		return nil
	}
	unix.Flock(int(f.fh.Fd()), unix.LOCK_EX)
	f.locked = true
	return nil
}
func (f *Flock) Close() error {
	if Reference {
		return f.f.Close()
	}
	f.m.Lock()
	defer f.m.Unlock()
	if !f.locked {
		return nil
	}
	unix.Flock(int(f.fh.Fd()), unix.LOCK_UN)
	f.locked = false
	f.fh.Close()
	f.fh = nil
	return nil
}
