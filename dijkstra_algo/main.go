package dijkstra_algo

import "math"

type Name string

type Cost int

type Neighbors map[Name]Cost

type Result map[Name]Cost

type LowestCost struct {
	Name Name
	Cost Cost
}

var (
	processed     []Name
	resultCosts   Result
	data          map[Name]Neighbors
	lowestCostTmp LowestCost
	startNode     Name
)

func GetPath(d map[Name]Neighbors, startNode Name) Result {
	initData(d, startNode)
	processNodeRecursive(&startNode)
	return resultCosts
}

func initData(d map[Name]Neighbors, startNode Name) {
	setStartNode(startNode)
	setData(d)
	resetProcessed()
	resetResultCosts()
	resetLowestCostTmp()
}

func setStartNode(s Name) {
	startNode = s
}

func resetProcessed() {
	processed = make([]Name, 0)
}

func setData(d map[Name]Neighbors) {
	data = d
}

func resetResultCosts() {
	resultCosts = make(map[Name]Cost)
	for node := range data {
		resultCosts[node] = getInf()
	}
}

func getInf() Cost {
	return math.MaxUint8
}

func isNodeProcessed(nodeName *Name) bool {
	for _, v := range processed {
		if v == *nodeName {
			return true
		}
	}
	return false
}

func processNodeRecursive(nodeName *Name) {
	if isNodeProcessed(nodeName) {
		return
	}
	processed = append(processed, *nodeName)
	setResultCosts(nodeName)
	setLowestCostTmp()
	processNodeRecursive(&lowestCostTmp.Name)

}

func getCostToCurrentNode(nodeName *Name) (c Cost) {
	if *nodeName == startNode {
		return 0
	}
	return resultCosts[*nodeName]
}

func setResultCosts(nodeName *Name) {
	c := getCostToCurrentNode(nodeName)
	neighbors := data[*nodeName]
	for neighborName, neighborCost := range neighbors {
		if !isNodeProcessed(&neighborName) {
			if (c + neighborCost) < resultCosts[neighborName] {
				resultCosts[neighborName] = c + neighborCost
			}
		}
	}
}

func resetLowestCostTmp() {
	lowestCostTmp.Cost = getInf()
	lowestCostTmp.Name = ``
}

func setLowestCostTmp() {
	prevLowestCostTmp := lowestCostTmp
	resetLowestCostTmp()
	for name, cost := range resultCosts {
		if !isNodeProcessed(&name) {
			if cost < lowestCostTmp.Cost {
				if prevLowestCostTmp.Cost != getInf() {
					lowestCostTmp.Cost = cost + prevLowestCostTmp.Cost
				} else {
					lowestCostTmp.Cost = cost
				}
				lowestCostTmp.Name = name
			}
		}
	}
	return
}
