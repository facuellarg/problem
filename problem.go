package problem

// import (
// 	"../estructuras/list"
// )

//Problem is a interface for every problem.
//this is a test
type Problem interface {
	IsGoal() bool
	Execute(action string) Problem
	PossibleActions() []string
	String() string
	Childrens() []Problem
}
