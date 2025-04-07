package main

import "fmt"

func main() {
	uf := NewDsu(6)

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(3, 4)

	fmt.Println(uf.IsConnected(0, 2)) // true
	fmt.Println(uf.IsConnected(0, 4)) // false

	uf.Union(2, 4)
	fmt.Println(uf.IsConnected(0, 4)) // true
}

type Dsu struct {
	Parent []int
	Rank   []int
}

func NewDsu(nodeCount int) *Dsu {
	unit := &Dsu{
		Parent: make([]int, nodeCount),
		Rank:   make([]int, nodeCount),
	}
	for i := 0; i < nodeCount; i++ {
		unit.Parent[i] = i
		unit.Rank[i] = 1
	}
	return unit
}

func (r *Dsu) Union(p, q int) {
	rootP := r.Find(p)
	rootQ := r.Find(q)
	if rootP == rootQ {
		return
	}
	if r.Rank[rootP] < r.Rank[rootQ] {
		r.Parent[rootP] = rootQ
	} else if r.Rank[rootP] > r.Rank[rootQ] {
		r.Parent[rootQ] = rootP
	} else {
		r.Parent[rootQ] = rootP
		r.Rank[rootP] += 1
	}
}

func (r *Dsu) IsConnected(p, q int) bool {
	if r.Find(p) == r.Find(q) {
		return true
	}
	return false
}

func (r *Dsu) Find(p int) int {
	if r.Parent[p] != p {
		r.Parent[p] = r.Find(r.Parent[p])
	}
	return r.Parent[p]
}
