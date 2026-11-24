package ds

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
