package breadth_first_search

/**
@param graph map[string][]string data, directed graph
@param start string vertex
@param destination string vertex
@return result int level
Example:
	if graph is:

		a -- b -- c -- f
		 |		   |
		 d -- i -- j

	start is "a"
	destination is "i"
	so, path should be ["a","d","i"] & result level should be 2

*/
func GetLevel(graph map[string][]string, start, destination string) (result int) {
	processed := map[string]bool{}
	processVertexRecursive(start, graph[start], destination, &graph, &result, &processed)
	return
}

func processVertexRecursive(start string, vertices []string, destination string, graph *map[string][]string, result *int, processed *map[string]bool) {
	*result++                              // recursive iterations counter
	nextLevelVertices := make([]string, 0) // queue for neighbours
	for _, v := range vertices {
		if v == start {
			continue // dont move in opposite direction
		}
		if (*processed)[v] {
			continue // optimization - dont move twice throw the one vertex
		}
		(*processed)[v] = true
		if v == destination {
			return // we found destination
		} else {
			nextLevelVertices = append(nextLevelVertices, (*graph)[v]...) // add all neighbours to queue
		}
	}
	processVertexRecursive(start, nextLevelVertices, destination, graph, result, processed) // process neighbours queue recursive
}
