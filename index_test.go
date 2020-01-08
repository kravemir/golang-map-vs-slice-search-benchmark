package golang_map_vs_slice_search_benchmark

import (
	"gotest.tools/assert"
	"testing"
)

type Index interface {
	Add(name string, value int64)
	Find(name string) (value int64, found bool)
}

var IndexElements = []string{
	"Name",
	"Description",
	"AddressLine1",
	"AddressLine2",
	"AddressLine3",
	"AddressZIP",
	"AddressCity",
	"AddressState",
	"AddressCountry",
	"Item01",
	"Item02",
	"Item03",
	"Item04",
	"FavouriteColour",
	"FavouriteWhiskeyBrand",
	"BirthDate",
	"Gender",

	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

var NonExistingElements = []string{
	"Name2",
	"AddressLine0",
	"AddressLine4",
	"Item00",
	"Item05",
	"Item06",
}

func IndexTest(t *testing.T, index Index) {
	for i, name := range IndexElements {
		index.Add(name, int64(i))
	}

	for i, name := range IndexElements {
		value, found := index.Find(name)

		assert.Equal(t, value, int64(i))
		assert.Equal(t, found, true)
	}

	for _, name := range NonExistingElements {
		_, found := index.Find(name)

		assert.Equal(t, found, false)
	}
}

func TestIntIndex(t *testing.T) {
	IndexTest(t, &IntIndex{})
}

func TestMapIndex(t *testing.T) {
	IndexTest(t, &MapIndex{})
}

func IndexBenachmark(b *testing.B, index Index) {
	for i, name := range IndexElements {
		index.Add(name, int64(i))
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i, name := range IndexElements {
			value, found := index.Find(name)

			if !found || value != int64(i) {
				panic("incorrect output")
			}
		}
	}

	for n := 0; n < b.N; n++ {
		for _, name := range NonExistingElements {
			_, found := index.Find(name)

			if found {
				panic("incorrect output")
			}
		}
	}
}

func BenchmarkFind(b *testing.B) {
	b.ReportAllocs()

	b.Run("IntIndex_Capacity_nil", func(b *testing.B) {
		IndexBenachmark(b, &IntIndex{})
	})

	b.Run("IntIndex_Capacity_100", func(b *testing.B) {
		IndexBenachmark(b, &IntIndex{
			names:  make([]string, 0, 100),
			values: make([]int64, 0, 100),
		})
	})

	b.Run("MapIndex", func(b *testing.B) {
		IndexBenachmark(b, &MapIndex{})
	})
}
