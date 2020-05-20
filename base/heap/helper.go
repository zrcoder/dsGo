package heap

type helper struct {
	slice []Any
	cmp   Cmp
}

func (h *helper) Len() int           { return len(h.slice) }
func (h *helper) Less(i, j int) bool { return h.cmp(i, j) }
func (h *helper) Swap(i, j int)      { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }
func (h *helper) Push(x interface{}) {
	h.slice = append(h.slice, x)
}
func (h *helper) Pop() interface{} {
	n := len(h.slice)
	x := h.slice[n-1]
	h.slice = h.slice[:n-1]
	return x
}
