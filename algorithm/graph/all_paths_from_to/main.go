package main

import (
	"fmt"
	"strings"
)

func main() {
	var node, edge int
	fmt.Scanf("%d %d", &node, &edge)
	graph := make([][]int, node+1)
	for i := 0; i < edge; i++ {
		var fromNode, toNode int
		fmt.Scanf("%d %d", &fromNode, &toNode)
		if len(graph[fromNode]) == 0 {
			graph[fromNode] = []int{}
		}
		graph[fromNode] = append(graph[fromNode], toNode)
	}
	fmt.Println(graph)

	resList := allArrivalPaths(graph)
	fmt.Println(resList)
	for _, row := range resList {
		var str = strings.Builder{}
		for j, v := range row {
			if len(row)-1 == j {
				str.WriteString(fmt.Sprintf("%d", v))
				fmt.Println(str.String())
			} else {
				str.WriteString(fmt.Sprintf("%d ", v))
			}
		}
	}
}

func allArrivalPaths(graph [][]int) [][]int {
	var resList = make([][]int, 0)
	var path = []int{}
	var existNode = make([]bool, len(graph)+1)
	path = append(path, 1)
	existNode[1] = true
	var backTracking func(graph [][]int, curNode int)
	backTracking = func(graph [][]int, curNode int) {
		if len(path) != 0 && path[len(path)-1] == len(graph)-1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			resList = append(resList, tmp)
			return
		}
		for _, v := range graph[curNode] {
			existNode[v] = true
			path = append(path, v)
			backTracking(graph, v)
			path = path[:len(path)-1]
			existNode[v] = false
		}
	}
	backTracking(graph, 1)
	return resList
}
