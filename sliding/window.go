package sliding

import (
	"container/list"
	"time"
)

type Window struct {
	maxRequests   int
	size          int64
	concurrentMap map[string]*list.List
}

func NewWindow(size int64, maxRequests int) *Window {
	return &Window{
		maxRequests: maxRequests,
		size:        size,
	}
}

func (w *Window) HandleRequest(ip string) bool {
	var clientTimeStamps *list.List = w.concurrentMap[ip]
	time := time.Now().UnixNano();

	for e := clientTimeStamps.Front(); e != nil; e = e.Next() {
		if time - e.Value.(int64) > w.size {
			clientTimeStamps.Remove(e)
		}
	}

	if clientTimeStamps.Len() < w.maxRequests {
		clientTimeStamps.PushBack(time)
		return true
	}

	return false;
}
