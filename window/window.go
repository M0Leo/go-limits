package window

type Window struct {
	time  int64
	count int
}

func NewWindow(time int64, count int) *Window {
	return &Window{
		time:  time,
		count: 0,
	}
}

func (w *Window) setRequestCount(count int) {
	w.count = count
}