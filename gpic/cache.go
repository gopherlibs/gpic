package gpic

import (
	"sync"
)

var cache inputCache = inputCache{items: make(map[string]int64)}

// A cache for gPic inputs. Stores them as a map and uses a RW mutex.
type inputCache struct {
	mu    sync.RWMutex
	items map[string]int64
}

func (this *inputCache) read(input string) (int64, bool) {

	this.mu.RLock()
	id, found := this.items[input]
	this.mu.RUnlock()

	if !found {
		return 0, false
	}

	return id, true
}

func (this *inputCache) write(input string, id int64) {

	this.mu.Lock()
	this.items[input] = id
	this.mu.Unlock()
}
