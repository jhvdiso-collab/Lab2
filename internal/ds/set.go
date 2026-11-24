package ds

type Set struct{ table Table[bool] }

func (s *Set) Add(k string)           { s.table.Put(k, true) }
func (s *Set) Delete(k string) bool   { return s.table.Delete(k) }
func (s *Set) Contains(k string) bool { _, ok := s.table.Get(k); return ok }
func (s *Set) Values() []string       { return s.table.Keys() }
