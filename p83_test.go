package euler

import "testing"

type p83Pair struct {
	x, y int
}

// FIFO queue (to be used as open list in BFS)
type p83Queue []p83Pair

func (q *p83Queue) Push(p p83Pair) {
	(*q) = append(*q, p)
}

func (q *p83Queue) Pop() p83Pair {
	p := (*q)[0]
	(*q) = (*q)[1:]
	return p
}

// Set of nodes with associated minimum path value (to be used as closed list)
type p83Set map[int]int64

func (s *p83Set) Get(p p83Pair) (int64, bool) {
	v, found := (*s)[(p.x<<16)+p.y]
	return v, found
}

func (s *p83Set) Set(p p83Pair, val int64) {
	(*s)[(p.x<<16)+p.y] = val
}

func TestP83(t *testing.T) {

	links := func(matrix [][]int64, sz int, cur p83Pair) []p83Pair {
		x, y := cur.x, cur.y
		nodes := make([]p83Pair, 0, 8)
		if y > 0 {
			nodes = append(nodes, p83Pair{x, y - 1})
		}
		if x > 0 {
			nodes = append(nodes, p83Pair{x - 1, y})
		}
		if x < sz-1 {
			nodes = append(nodes, p83Pair{x + 1, y})
		}
		if y < sz-1 {
			nodes = append(nodes, p83Pair{x, y + 1})
		}
		return nodes
	}

	solve := func(matrix [][]int64, sz int) int64 {
		// Standard BFS search, with open/closed sets
		q := p83Queue{} // open
		s := p83Set{}   // closed
		q.Push(p83Pair{0, 0})
		s.Set(p83Pair{0, 0}, matrix[0][0])

		for len(q) > 0 {
			cur := q.Pop()
			curval, _ := s.Get(cur)
			for _, next := range links(matrix, sz, cur) {
				nextval := curval + matrix[next.y][next.x]
				if v, found := s.Get(next); found {
					if v > nextval {
						q.Push(next)
						s.Set(next, nextval)
					}
				} else {
					q.Push(next)
					s.Set(next, nextval)
				}
			}
		}

		sol, _ := s.Get(p83Pair{sz - 1, sz - 1})
		return sol
	}

	matrix := p81TestMatrix()
	sol := solve(matrix, 5)
	if sol != 2297 {
		t.Error(sol)
	}

	matrix = p81ParseMatrix(p81Matrix)
	sol = solve(matrix, 80)
	if sol != 2297 {
		t.Error(sol)
	}

}
