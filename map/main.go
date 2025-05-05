package main

import "fmt"

func main() {
	// Initialize the map with some values
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	if _, ok := m[1]; ok {
		fmt.Println("Key 1 exists in the map")
	} else {
		fmt.Println("Key 1 does not exist in the map")
	}

	a := make(map[int]int)
	a[1] = 1
	a[2] = 2

	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
	}
	fmt.Println(graph[1]) // Output: [2 3]
	fmt.Println(graph[2]) // Output: [4 5]
	fmt.Println(graph[3]) // Output: [6]

	dsf(graph, 1, make(map[int]bool))
}

func dsf(graph map[int][]int, vertex int, visited map[int]bool) {
	visited[vertex] = true
	fmt.Printf("-> %d", vertex)

	for _, v := range graph[vertex] {
		if !visited[v] {
			dsf(graph, v, visited)
		}
	}

}
