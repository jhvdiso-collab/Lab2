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

type Set struct{ table Table[bool] }

func (s *Set) Add(k string)           { s.table.Put(k, true) }
func (s *Set) Delete(k string) bool   { return s.table.Delete(k) }
func (s *Set) Contains(k string) bool { _, ok := s.table.Get(k); return ok }
func (s *Set) Values() []string       { return s.table.Keys() }

type Node struct {
	Value       int
	Left, Right *Node
}
type Tree struct{ Root *Node }

func (t *Tree) Insert(v int) (int, bool) {
	p := &t.Root
	depth := 1
	for *p != nil {
		if v == (*p).Value {
			return 0, false
		}
		if v < (*p).Value {
			p = &(*p).Left
		} else {
			p = &(*p).Right
		}
		depth++
	}
	*p = &Node{Value: v}
	return depth, true
}

type Bloom struct{ bits []uint64 }

func (b *Bloom) Add(k string) {
	if len(b.bits) == 0 {
		b.bits = make([]uint64, 1024)
	}
	for i := uint64(0); i < 7; i++ {
		n := Hash(k, i) % uint64(len(b.bits)*64)
		b.bits[n/64] |= 1 << (n % 64)
	}
}
func (b *Bloom) Maybe(k string) bool {
	if len(b.bits) == 0 {
		return false
	}
	for i := uint64(0); i < 7; i++ {
		n := Hash(k, i) % uint64(len(b.bits)*64)
		if b.bits[n/64]&(1<<(n%64)) == 0 {
			return false
		}
	}
	return true
}

type FilteredTable struct {
	Table  Table[string]
	Filter Bloom
}

func (t *FilteredTable) Put(k, v string) { t.Filter.Add(k); t.Table.Put(k, v) }
func (t *FilteredTable) Get(k string) (string, bool) {
	if !t.Filter.Maybe(k) {
		return "", false
	}
	return t.Table.Get(k)
}
func (t *FilteredTable) Delete(k string) bool { return t.Table.Delete(k) }

// LFU breaks frequency ties by least recent access.
type cacheEntry struct {
	value, freq int
	time        uint64
}
type LFU struct {
	table    Table[cacheEntry]
	capacity int
	clock    uint64
}

func NewLFU(capacity int) *LFU { return &LFU{capacity: capacity} }
func (c *LFU) Get(k string) (int, bool) {
	e, ok := c.table.Get(k)
	if !ok {
		return 0, false
	}
	c.clock++
	e.freq++
	e.time = c.clock
	c.table.Put(k, e)
	return e.value, true
}
func (c *LFU) Set(k string, v int) {
	if c.capacity <= 0 {
		return
	}
	c.clock++
	if e, ok := c.table.Get(k); ok {
		e.value = v
		e.freq++
		e.time = c.clock
		c.table.Put(k, e)
		return
	}
	if c.table.Len() == c.capacity {
		var victim string
		var best cacheEntry
		first := true
		for _, key := range c.table.Keys() {
			e, _ := c.table.Get(key)
			if first || e.freq < best.freq || (e.freq == best.freq && e.time < best.time) {
				victim, best, first = key, e, false
			}
		}
		c.table.Delete(victim)
	}
	c.table.Put(k, cacheEntry{v, 1, c.clock})
}
