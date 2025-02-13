package genericlist

type Genericlist[T comparable] struct {
	data []T
}

func New[T comparable]() *Genericlist[T] {
	return &Genericlist[T]{
		data: []T{},
	}
}

func (l *Genericlist[T]) Insert(item T) {
	l.data = append(l.data, item)
}

func (l *Genericlist[T]) Get(i int) T {
	if i > len(l.data)-1 {
		panic("index out of range")
	}

	for it := 0; i < len(l.data); i++ {
		if i == it {
			return l.data[it]
		}
	}

	panic("value not found")
}

func (l *Genericlist[T]) RemoveByValue(item T) {
	for it := 0; it < len(l.data); it++ {
		if l.data[it] == item {
			l.data = append(l.data[:it], l.data[it+1:]...)
		}
	}
}

func (l *Genericlist[T]) Remove(i int) {
	if i > len(l.data)-1 {
		panic("index out of range")
	}

	for it := 0; it < len(l.data); it++ {
		if i == it {
			l.data = append(l.data[:it], l.data[it+1:]...)
		}
	}
}
