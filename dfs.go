package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

// enums for dfs vertice states
const (
	WHITE     = iota
	GRAY      = iota
	BLACK     = iota
	UNDEFINED = -1
)

// getColor returns the string representing the color
func getColor(color int) string {
	switch color {
	case WHITE:
		return "white"
	case GRAY:
		return "gray"
	case BLACK:
		return "black"
	}
	return "error"
}

type vertex struct {
	name      int
	neighbors []int
	color     int
	parent    int
}

func (vertex *vertex) addNeighbor(name int) {
	for _, v := range vertex.neighbors {
		if v == name {
			fmt.Println("already in")
			return
		}
	}
	vertex.neighbors = append(vertex.neighbors, name)
	sort.Ints(vertex.neighbors)
}

func (vertex vertex) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(strconv.Itoa(vertex.name))
	buffer.WriteString("| color: ")
	buffer.WriteString(getColor(vertex.color))
	buffer.WriteString(" parent: ")
	buffer.WriteString(strconv.Itoa(vertex.parent))
	buffer.WriteString(" neighbors: {")
	for i, v := range vertex.neighbors {
		if i != 0 && i != len(vertex.neighbors) {
			buffer.WriteString(",")
		}
		buffer.WriteString(strconv.Itoa(v))
	}
	buffer.WriteString("}]")
	return buffer.String()
}

// Graph represents a graph
type Graph struct {
	vertices []vertex
}

func (graph Graph) Len() int {
	return len(graph.vertices)
}

func (graph Graph) Swap(i, j int) {
	graph.vertices[i], graph.vertices[j] = graph.vertices[j], graph.vertices[i]
}

func (graph Graph) Less(i, j int) bool {
	return graph.vertices[i].name < graph.vertices[j].name
}

func (graph *Graph) initializeGraph() {
	for i := 0; i < len(graph.vertices); i++ {
		graph.vertices[i].color = WHITE
		graph.vertices[i].parent = UNDEFINED
	}
}

// AddEdge adds an edge to the graph
func (graph *Graph) AddEdge(a int, b int) {
	// first look if vertex exists
	for i, vertex := range graph.vertices {
		if vertex.name == a {
			graph.vertices[i].addNeighbor(b)
			return
		}
	}

	// if we make it here, the vertex doesnt exist
	graph.vertices = append(graph.vertices, vertex{
		name:      a,
		color:     WHITE,
		neighbors: []int{b},
	})
	sort.Sort(*graph)
}

func (graph Graph) String() string {
	var buffer bytes.Buffer
	for i, vertex := range graph.vertices {
		buffer.WriteString(vertex.String())
		if i != len(graph.vertices)-1 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

// PrintGraph prints out the graph
func (graph Graph) PrintGraph() {
	fmt.Println(graph.String())
}

// LoadGraph loads a graph to have dfs run
func LoadGraph() *Graph {
	g := new(Graph)
	return g
}

func main() {
	g := LoadGraph()
	g.AddEdge(1, 3)
	g.AddEdge(4, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 1)
	g.initializeGraph()
	g.PrintGraph()
}
