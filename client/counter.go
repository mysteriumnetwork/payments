package client

import (
	"runtime"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

var Counter = callCounter{}

type callCounter struct {
	stats map[string]uint64
	lock  sync.Mutex
	once  sync.Once
}

func (cc *callCounter) print() {
	spew.Dump(cc.stats)
}

func (cc *callCounter) Count() {
	cc.lock.Lock()
	defer cc.lock.Unlock()
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		cc.once.Do(func() {
			cc.stats = make(map[string]uint64)
		})

		name := details.Name()
		if v, ok := cc.stats[name]; ok {
			cc.stats[name] = v + 1
		} else {
			cc.stats[name] = 1
		}

		cc.print()
	}
}
