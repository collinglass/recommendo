package sort

import (
	"github.com/collinglass/recommendo/algo"
	"github.com/collinglass/recommendo/data"
	"sort"
)

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []data.Similar

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Distance > p[j].Distance }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[int]data.Similar) PairList {
	p := make(PairList, len(m))
	i := 0
	for _, v := range m {
		p[i] = v
		i++
	}
	sort.Sort(p)
	return p
}

// Implement Algo and get Top Matches
/* n is size of list, person is data.UserId,
users is map of data.User and
simFunc is algo.SimFunc */
func TopMatches(n int, person int, users map[int]data.User, simFunc algo.SimFunc) PairList {
	algo.GetSimilar(users, simFunc)
	sort := sortMapByValue(users[person].Similars)

	return sort[0:n]
}
