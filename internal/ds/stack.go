package ds

type link[T any] struct {
	value T
	next  *link[T]
}
type Stack[T any] struct{ head *link[T] }

func (s *Stack[T]) Push(v T) { s.head = &link[T]{v, s.head} }
func (s *Stack[T]) Pop() (T, bool) {
	if s.head == nil {
		var z T
		return z, false
	}
	v := s.head.value
	s.head = s.head.next
	return v, true
}
