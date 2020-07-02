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
	processVertexRecursive(graph[start], destination, &graph, &result, &processed)
	return
}

func processVertexRecursive(vertices []string, destination string, graph *map[string][]string, result *int, processed *map[string]bool) {
	*result++
	nextLevelVertices := make([]string, 0)
	for _, v := range vertices {
		if (*processed)[v] {
			continue
		}
		(*processed)[v] = true
		if v == destination {
			return
		} else {
			nextLevelVertices = append(nextLevelVertices, (*graph)[v]...)
		}
	}
	processVertexRecursive(nextLevelVertices, destination, graph, result, processed)
}
