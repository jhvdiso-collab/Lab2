package ds

// Hash mixes a string with a seed; it is also used by the Bloom filter.
func Hash(s string, seed uint64) uint64 {
	h := uint64(14695981039346656037) ^ seed
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211
	}
	h ^= h >> 33
	h *= 0xff51afd7ed558ccd
	h ^= h >> 33
	return h
}

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

type Cuckoo[V any] struct {
	slots [2][]*entry[V]
	seed  uint64
	size  int
}

func (c *Cuckoo[V]) Get(k string) (V, bool) {
	if len(c.slots[0]) > 0 {
		for j := 0; j < 2; j++ {
			p := c.slots[j][Hash(k, c.seed+uint64(j))%uint64(len(c.slots[j]))]
			if p != nil && p.key == k {
				return p.value, true
			}
		}
	}
	var z V
	return z, false
}
func (c *Cuckoo[V]) place(p *entry[V]) *entry[V] {
	j := 0
	for n := 0; n < 4*len(c.slots[0]); n++ {
		i := Hash(p.key, c.seed+uint64(j)) % uint64(len(c.slots[j]))
		p, c.slots[j][i] = c.slots[j][i], p
		if p == nil {
			return nil
		}
		j = 1 - j
	}
	return p
}
func (c *Cuckoo[V]) Put(k string, v V) {
	if len(c.slots[0]) == 0 {
		c.slots[0] = make([]*entry[V], 16)
		c.slots[1] = make([]*entry[V], 16)
	}
	for j := 0; j < 2; j++ {
		p := c.slots[j][Hash(k, c.seed+uint64(j))%uint64(len(c.slots[j]))]
		if p != nil && p.key == k {
			p.value = v
			return
		}
	}
	c.size++
	p := c.place(&entry[V]{key: k, value: v})
	if p == nil && c.size <= len(c.slots[0]) {
		return
	}
	all := []*entry[V]{}
	if p != nil {
		all = append(all, p)
	}
	for _, a := range c.slots {
		for _, e := range a {
			if e != nil {
				all = append(all, e)
			}
		}
	}
	n := len(c.slots[0]) * 2
	for {
		c.seed += 2
		c.slots[0] = make([]*entry[V], n)
		c.slots[1] = make([]*entry[V], n)
		ok := true
		for _, e := range all {
			if c.place(e) != nil {
				ok = false
				break
			}
		}
		if ok {
			return
		}
		n *= 2
	}
}
func (c *Cuckoo[V]) Delete(k string) bool {
	if len(c.slots[0]) == 0 {
		return false
	}
	for j := 0; j < 2; j++ {
		i := Hash(k, c.seed+uint64(j)) % uint64(len(c.slots[j]))
		if p := c.slots[j][i]; p != nil && p.key == k {
			c.slots[j][i] = nil
			c.size--
			return true
		}
	}
	return false
}
