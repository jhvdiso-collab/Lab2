package ds

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
