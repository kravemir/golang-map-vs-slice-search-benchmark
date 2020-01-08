package golang_map_vs_slice_search_benchmark

type MapIndex map[string]int64

func (m MapIndex) Add(name string, value int64) {
	m[name] = value
}

func (m MapIndex) Find(name string) (value int64, found bool) {
	value, found = m[name]
	return
}
