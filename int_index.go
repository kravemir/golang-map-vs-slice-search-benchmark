package golang_map_vs_slice_search_benchmark

import "sort"

type IntIndex struct {
	names  []string
	values []int64
}

func (index *IntIndex) Add(name string, value int64) {
	index.names = append(index.names, name)
	index.values = append(index.values, value)

	sort.Sort((*intIndexSorter)(index))
}

func (index *IntIndex) Find(name string) (value int64, found bool) {
	i := sort.SearchStrings(index.names, name)

	if i < len(index.names) && index.names[i] == name {
		return index.values[i], true
	}

	return 0, false
}

type intIndexSorter IntIndex

func (index *intIndexSorter) Len() int {
	return len(index.names)
}

func (index *intIndexSorter) Less(i, j int) bool {
	return index.names[i] < index.names[j]
}

func (index *intIndexSorter) Swap(i, j int) {
	index.names[i], index.names[j] = index.names[j], index.names[i]
	index.values[i], index.values[j] = index.values[j], index.values[i]
}
