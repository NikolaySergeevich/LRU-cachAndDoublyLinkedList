package list

type Key interface{}

type ListItem struct {
	Value interface{}
	Key   interface{}
	Next *ListItem
	Prev *ListItem
}

type List struct {
	front *ListItem
	back  *ListItem
	len   int
}

// Получение длины списка
func (l *List) Len() int {
	return l.len
}

// Получение первого элемента списка
func (l *List) Front() *ListItem {
	return l.front
}

// Получение послежнего элемента списка
func (l *List) Back() *ListItem {
	return l.back
}

// Добавление нового значения в начало списка
func (l *List) PushFront(k, v interface{}) *ListItem {
	newItem := &ListItem{Key: k, Value: v}
	if l.len == 0 {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Next = l.front
		l.front.Prev = newItem
		l.front = newItem
	}
	l.len++
	return newItem
}

// Добавление нового значения в конец списка
func (l *List) PushBack(k, v interface{}) *ListItem {
	newItem := &ListItem{Key: k, Value: v}
	if l.len == 0 {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Prev = l.back
		l.back.Next = newItem
		l.back = newItem
	}
	l.len++
	return newItem
}

// Удаление элемента из списка
func (l *List) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	l.len--
}

// Перемещение элемента в начало списка
func (l *List) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}
	l.Remove(i)
	i.Prev = nil
	i.Next = l.front
	if l.front != nil {
		l.front.Prev = i
	}
	l.front = i
	if l.len == 0 {
		l.back = i
	}
	l.len++
}

type Cache struct {
	capacity int
	Queue    *List
	Items    map[Key]*ListItem
}

// Создание нового кэша
func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		Queue:    &List{},
		Items:    make(map[Key]*ListItem),
	}
}

func (c *Cache) Set(key Key, value interface{}) bool {
	if item, ok := c.Items[key]; ok {
		item.Value = value
		c.Queue.MoveToFront(item)
		return true
	}
	if c.Queue.Len() == c.capacity {
		last := c.Queue.Back()
		c.Queue.Remove(last)
		delete(c.Items, last.Key)
	}
	item := c.Queue.PushFront(key, value)
	c.Items[key] = item
	return false
}

func (c *Cache) Get(key Key) (interface{}, bool) {
	if item, ok := c.Items[key]; ok {
		c.Queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (c *Cache) Clear() {
	c.Queue = &List{}
	c.Items = make(map[Key]*ListItem)
}
