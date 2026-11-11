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
