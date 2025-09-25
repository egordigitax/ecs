package ecs

type Storage[T any] struct {
	data map[ID]any
}

func (s *Storage[T]) Add(id ID, c T) {
	s.data[id] = &c
}

func (s *Storage[T]) Get(id ID) (*T, bool) {
	if v, ok := s.data[id]; ok {
		return v.(*T), true
	}
	return nil, false
}

func (s *Storage[T]) Remove(id ID) {
	delete(s.data, id)
}

func (s *Storage[T]) Each(fn func(id ID, c *T)) {
	for id, val := range s.data {
		fn(id, val.(*T))
	}
}

func (s *Storage[T]) Raw() map[ID]any {
	return s.data
}

func GetStorage[T any](w *World, ct ComponentType) *Storage[T] {
	if s, ok := w.storages[ct]; ok {
		return s.(*Storage[T])
	}
	return NewStorage[T](w, ct)
}

func NewStorage[T any](w *World, ct ComponentType) *Storage[T] {
	s := &Storage[T]{data: make(map[ID]any)}
	w.storages[ct] = s
	return s
}
