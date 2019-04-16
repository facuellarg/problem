package problem

import (
	"container/list"
	"fmt"
	"github.com/facuellarg/fronteir"
)

//FAIL print fail if no  have solution
const FAIL = "No hay camino"

//SearchProblem struc
type SearchProblem struct {
	init Problem
}

//NewSearchProblem give a new search problem.
func NewSearchProblem(p Problem) (sh SearchProblem) {
	sh.init = p
	return sh
}

func searchProblem(init Problem, fronteir fronteir.Fronteir) (string, int) {
	explored := make(map[string]bool)
	nodosExpandidos := 0
	n := NewNode(init, "")
	if n.Problem().IsGoal() {
		return n.Path(), nodosExpandidos
	}
	fronteir.Add(n)
	for fronteir.Size() != 0 {
		n = fronteir.Pop().(Node)
		nodosExpandidos++
		explored[n.Problem().String()] = true
		for _, action := range n.Problem().PossibleActions() {
			p := n.Problem().Execute(action)
			if !explored[p.String()] {
				if p.IsGoal() {
					return n.Path() + action + " ", nodosExpandidos
				}
				fronteir.Add(NewNode(p, n.Path()+action+" "))

			}
		}

	}
	return FAIL, nodosExpandidos
}

//BFS use queue in search general problem tu solved
func (sh SearchProblem) BFS() (string, int) {
	var queue fronteir.Queue
	return searchProblem(sh.init, &queue)
}

//DFS use stack in search general problem tu solved
func (sh SearchProblem) DFS() (string, int) {
	var stack fronteir.Stack
	return searchProblem(sh.init, &stack)
}

//LDFS Limited dfs
func LDFS(init Problem, limit int, explored map[string]int) (string, int) {
	n := NewNode(init, "")
	if n.Problem().IsGoal() {
		return n.Path(), 0
	}
	nodosExpandidos := 0
	currentDepth := 0
	fronteir := list.New()
	depth := list.New()
	isInFronteir := make(map[string]int)
	fronteir.PushBack(n)
	depth.PushBack(currentDepth)
	isInFronteir[n.Problem().String()] = currentDepth
	for fronteir.Len() != 0 {
		n = fronteir.Remove(fronteir.Back()).(Node)
		currentDepth = depth.Remove(depth.Back()).(int)
		nodosExpandidos++
		delete(isInFronteir, n.Problem().String())
		explored[n.Problem().String()] = currentDepth
		if currentDepth < limit {
			for _, action := range n.Problem().PossibleActions() {
				child := n.Problem().Execute(action)
				valueExplored, isExplored := explored[child.String()]
				valueFronteir, isInF := isInFronteir[child.String()]
				if !isExplored || (isExplored && valueExplored >= currentDepth+1) {
					if !isInF || (isInF && valueFronteir > currentDepth+1) {
						if child.IsGoal() {
							return n.Path() + action, nodosExpandidos
						}
						fronteir.PushBack(NewNode(child, n.Path()+action+" "))
						depth.PushBack(currentDepth + 1)
						isInFronteir[child.String()] = currentDepth + 1
						continue
					}
				}
			}
		}
	}

	return FAIL, nodosExpandidos

}

//IDFS usa DLS de forma iterativa para resolver el problema.
func (sh SearchProblem) IDFS() (result string, nodos int) {
	fmt.Print()
	explored := make(map[string]int)
	for i := 0; i < 40; i++ {
		result, nodos = LDFS(sh.init, i, explored)
		if result != FAIL {
			fmt.Println("Altura: ", i)
			return result, nodos
		}
	}
	result = FAIL
	return
}
