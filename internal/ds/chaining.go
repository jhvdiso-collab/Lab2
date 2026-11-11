package ds

type entry[V any] struct {
	key   string
	value V
	next  *entry[V]
}
type Table[V any] struct {
	buckets []*entry[V]
	size    int
}

func (t *Table[V]) Put(k string, v V) {
	if len(t.buckets) == 0 {
		t.buckets = make([]*entry[V], 16)
	}
	i := Hash(k, 0) % uint64(len(t.buckets))
	for p := t.buckets[i]; p != nil; p = p.next {
		if p.key == k {
			p.value = v
			return
		}
	}
	t.buckets[i] = &entry[V]{k, v, t.buckets[i]}
	t.size++
	if t.size > len(t.buckets)*2 {
		old := t.buckets
		t.buckets = make([]*entry[V], len(old)*2)
		t.size = 0
		for _, p := range old {
			for ; p != nil; p = p.next {
				t.Put(p.key, p.value)
			}
		}
	}
}
func (t *Table[V]) Get(k string) (V, bool) {
	if len(t.buckets) > 0 {
		for p := t.buckets[Hash(k, 0)%uint64(len(t.buckets))]; p != nil; p = p.next {
			if p.key == k {
				return p.value, true
			}
		}
	}
	var z V
	return z, false
}
func (t *Table[V]) Delete(k string) bool {
	if len(t.buckets) == 0 {
		return false
	}
	p := &t.buckets[Hash(k, 0)%uint64(len(t.buckets))]
	for *p != nil {
		if (*p).key == k {
			*p = (*p).next
			t.size--
			return true
		}
		p = &(*p).next
	}
	return false
}
func (t *Table[V]) Keys() []string {
	r := []string{}
	for _, p := range t.buckets {
		for ; p != nil; p = p.next {
			r = append(r, p.key)
		}
	}
	return r
}
func (t *Table[V]) Len() int { return t.size }
