package problem

// import (
// 	"../estructuras/list"
// )

//Problem is a interface for every problem.
//this is a test
type Problem interface {
	IsGoal() bool
	Execute(action string) (Problem, int)
	PossibleActions() []string
	String() string
	State() interface{}
}
