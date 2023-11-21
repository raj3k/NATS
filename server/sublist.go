package server

// type SublistResult struct {
// 	psubs []*subscription
// }

type Sublist struct {
	slr map[string][]*subscription
}

func NewSublist() *Sublist {
	return &Sublist{
		slr: make(map[string][]*subscription),
	}
}
