package ds

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
