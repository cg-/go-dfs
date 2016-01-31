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
	start     int
	finish    int
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
	buffer.WriteString(" start: ")
	buffer.WriteString(strconv.Itoa(vertex.start))
	buffer.WriteString(" finish: ")
	buffer.WriteString(strconv.Itoa(vertex.finish))
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
		graph.vertices[i].start = UNDEFINED
		graph.vertices[i].finish = UNDEFINED
		graph.vertices[i].parent = UNDEFINED
	}
}

func (graph *Graph) vertexExists(name int) bool {
	for _, vertex := range graph.vertices {
		if vertex.name == name {
			return true
		}
	}
	return false
}

func (graph *Graph) addVertex(name int) {
	graph.vertices = append(graph.vertices, vertex{
		name: name,
	})

	sort.Sort(*graph)
}

func (graph *Graph) findVertex(a int) *vertex {
	for i, v := range graph.vertices {
		if v.name == a {
			return &graph.vertices[i]
		}
	}
	fmt.Println("error in findVertex")
	return &vertex{}
}

// AddEdge adds an edge to the graph
func (graph *Graph) AddEdge(a int, b int) {
	// first look if vertex exists
	if !graph.vertexExists(a) {
		graph.addVertex(a)
	}
	if !graph.vertexExists(b) {
		graph.addVertex(b)
	}
	graph.findVertex(a).addNeighbor(b)
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

func (graph *Graph) dfs() {
	time := 0
	// run through all vertices on the graph
	for i := 0; i < graph.Len(); i++ {
		if graph.vertices[i].color == WHITE {
			graph.visit(graph.vertices[i].name, &time)
		}
	}
}

func (graph *Graph) visit(vertex int, time *int) {
	curVer := graph.findVertex(vertex)
	curVer.color = GRAY
	*time = *time + 1
	curVer.start = *time

	for i, v := range curVer.neighbors {
		if graph.findVertex(v).color != WHITE {
			continue
		}
		graph.findVertex(v).parent = vertex
		graph.visit(curVer.neighbors[i], time)
	}
	curVer.color = BLACK
	*time = *time + 1
	curVer.finish = *time
}

func main() {
	g := LoadGraph()
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(2, 3)
	g.AddEdge(2, 1)
	g.AddEdge(6, 7)
	fmt.Println("\nBefore DFS")
	g.PrintGraph()
	fmt.Println("\nAfter DFS")
	g.dfs()
	g.PrintGraph()
}
