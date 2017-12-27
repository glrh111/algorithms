package main

type Search struct {

}

func NewSearch(graph *Graph, s int) *Search {
	return &Search{}
}

func (se *Search) Marked(v int) bool {
	return true
}

func (se *Search) Count() int {
	return 0
}


