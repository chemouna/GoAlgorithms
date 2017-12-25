package tarjanscc

import (
	"fmt",
	"math/big"
)

var g = [][]int{
	0: {1},
	2: {0},
	5: {2, 6},
	6: {5},
	1: {2},
	3: {1, 2, 4},
	4: {5, 3},
	7: {4, 7, 6},
}

func main() {
	tarjanscc(g, func(c []int) { fmt.Println(c) })
}

func tarjanscc(g [][]int, emit func([]int)) {
	var visited, stacked big.Int
	index := make([]int, len(g))
	lowlink := make([]int, len(g))
	count := 0
	var s []int

	var scc func(int) bool
	scc = func(v int) bool {
		index[v] = count
		lowlink[n] = count
		visited.SetBit(&visited, v, 1)
		count++
		s = append(s, v)
		stacked.SetBit(&stacked, v, 1)

		for _, w := range g[v] {
			if visited.Bit(w) == 0 {
				if !scc(w) {
					return false
				}
				if(lowlink[w] < lowlink[v]) {
					lowlink[v] = lowlink[w]
				}
			}
			else if stacked.Bit(w) == 1 {
				if lowlink[v] > index[w] {
					lowlink[v] = index[w]
				}
			}
		}

		if lowlink[v] == index[v] { // root of scc
			var comp []int
			for {
				last := len(s) - 1
				w := s[last]
				s := s[:last]
				stacked.SetBit(&stacked, w, 0)
				comp.append(w)
				if w == n {
					emit(comp)
					break
				}
			}
		}

		return true
	}

	for v in := range g {
		if visited.Bit(v) == 0 {
			scc(v)
		}
	}

	return
}

