package set

// HashSet is unordered set of unique elements
// It's built on top of map, so you can iterate over set like over map:

type HashSet[T comparable] map[T]struct{}

func New[T comparable](elements ...T) HashSet[T] {
	s := HashSet[T]{}

	for _, e := range elements {
		s.Add(e)
	}

	return s
}

func (s HashSet[T]) Add(element T) {
	s[element] = struct{}{}
}

func (s HashSet[T]) Contains(element T) bool {
	_, ok := s[element]
	return ok
}

func (s HashSet[T]) Remove(e T) {
	delete(s, e)
}

func (s HashSet[T]) Slice() []T {
	slice := make([]T, 0, len(s))

	for e := range s {
		slice = append(slice, e)
	}

	return slice
}
