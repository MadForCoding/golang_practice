package main

import (
	"fmt"
)

// input
//5 4
//1 2
//1 3
//2 4
//3 4
//1 4

// output 1

func main() {
	var points, edges int
	var source, dest int
	fmt.Scanf("%d %d", &points, &edges)
	dsu := NewDsu(points)
	for i := 0; i < edges; i++ {
		var fromPoint, toPoint int
		fmt.Scanf("%d %d", &fromPoint, &toPoint)
		dsu.union(fromPoint, toPoint)
	}
	fmt.Scanf("%d %d", &source, &dest)

	if dsu.find(source) == dsu.find(dest) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

type Dsu struct {
	parent []int
	rank   []int
}

func NewDsu(point int) *Dsu {
	res := &Dsu{
		parent: make([]int, point+1),
		rank:   make([]int, point+1),
	}
	for i := 1; i <= point; i++ {
		res.parent[i] = i
		res.rank[i] = 1
	}
	return res
}

func (r *Dsu) find(point int) int {
	if r.parent[point] != point {
		r.parent[point] = r.find(r.parent[point])
	}
	return r.parent[point]
}

func (r *Dsu) union(x, y int) {
	rootX := r.find(x)
	rootY := r.find(y)
	if r.rank[rootX] < r.rank[rootY] {
		r.parent[rootX] = rootY
	} else if r.rank[rootX] > r.rank[rootY] {
		r.parent[rootY] = rootX
	} else {
		r.parent[rootX] = rootY
		r.rank[rootY] += 1
	}
}
