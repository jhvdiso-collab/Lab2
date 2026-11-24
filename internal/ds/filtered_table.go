package ds

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
