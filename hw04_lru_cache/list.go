package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len       int
	firstNode *ListItem
	lastNode  *ListItem
}

func (l *list) Remove(i *ListItem) {
	if l.len == 0 {
		return
	}
	if i.Prev == nil {
		l.firstNode = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.lastNode = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.firstNode
}

func (l list) Back() *ListItem {
	return l.lastNode
}

func (l *list) PushFront(v interface{}) *ListItem {
	fn := ListItem{v, l.firstNode, nil}
	if fn.Next != nil {
		fn.Next.Prev = &fn
	} else {
		l.lastNode = &fn
	}

	l.firstNode = &fn
	l.len++
	return l.firstNode
}

func (l *list) PushBack(v interface{}) *ListItem {
	ln := ListItem{v, nil, l.lastNode}
	if ln.Prev != nil {
		ln.Prev.Next = &ln
	}
	l.lastNode = &ln
	l.len++
	return l.lastNode
}

func NewList() List {
	l := new(list)
	return l
}
