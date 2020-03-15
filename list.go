package leetcode

// SliceList is a stack queue
type SliceList []interface{}

// Len returns length
func (l *SliceList) Len() int { return len(*l) }

// Push method
func (l *SliceList) Push(v ...interface{}) {
	*l = append(*l, v...)
}

// Pop method
func (l *SliceList) Pop() interface{} {
	length := l.Len()
	if length == 0 {
		return nil
	}

	v := (*l)[length-1]
	*l = (*l)[:length-1]
	return v
}

// Unshift method
func (l *SliceList) Unshift(v interface{}) {
	if l.Len() == 0 {
		*l = append(*l, v)
	} else {
		*l = append((*l)[:1], *l)
		(*l)[0] = v
	}
}

// Shift method
func (l *SliceList) Shift() interface{} {
	length := l.Len()
	if length == 0 {
		return nil
	}

	v := (*l)[0]
	*l = (*l)[1:]
	return v
}
