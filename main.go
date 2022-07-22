package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Graph : represents a Graph
type Graph struct {
	nodes []*GraphNode
}

// GraphNode : represents a Graph node
type GraphNode struct {
	userId         int
	communications map[int]int
}

// New : returns a new instance of a Graph
func New() *Graph {
	return &Graph{
		nodes: []*GraphNode{},
	}
}

// AddNode : adds a new node to the Graph
func (g *Graph) AddNode() (id int) {
	id = len(g.nodes)
	g.nodes = append(g.nodes, &GraphNode{
		userId:         id,
		communications: make(map[int]int),
	})
	return
}

// AddEdge : adds a directional edge together with a weight
func (g *Graph) AddEdge(n1, n2 int, w int) {
	g.nodes[n1].communications[n2] = w
}

// Neighbors : returns a list of node IDs that are linked to this node
func (g *Graph) Neighbors(id int) []int {
	neighbors := []int{}
	for _, node := range g.nodes {
		for edge := range node.communications {
			if node.userId == id {
				neighbors = append(neighbors, edge)
			}
			if edge == id {
				neighbors = append(neighbors, node.userId)
			}
		}
	}
	return neighbors
}

// Nodes : returns a list of node IDs
func (g *Graph) Nodes() []int {
	nodes := make([]int, len(g.nodes))
	for i := range g.nodes {
		nodes[i] = i
	}
	return nodes
}

// communications : returns a list of communications with weights
func (g *Graph) communications() [][3]int {
	communications := make([][3]int, 0, len(g.nodes))
	for i := 0; i < len(g.nodes); i++ {
		for k, v := range g.nodes[i].communications {
			communications = append(communications, [3]int{i, k, int(v)})
		}
	}
	return communications
}

func main() {

	graph := New()
	user0 := graph.AddNode()
	user1 := graph.AddNode()
	user2 := graph.AddNode()
	user3 := graph.AddNode()

	graph.AddEdge(user0, user1, 1)
	graph.AddEdge(user1, user2, 5)
	graph.AddEdge(user0, user3, 1)
	graph.AddEdge(user3, user1, 4)

	r := mux.NewRouter()
	r.HandleFunc("/graph", graph.getAllEdges).Methods("GET")
	r.HandleFunc("/graph/{id}", graph.getEdgeById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func (g *Graph) getAllEdges(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g.communications())
}

func (g *Graph) getEdgeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	for i := 1; i < len(g.communications()); i++ {
		fmt.Println(g.communications()[i])
		s := strconv.Itoa(i)
		if s == params["id"] {
			json.NewEncoder(w).Encode(g.communications()[i])
			return
		}
	}
	json.NewEncoder(w).Encode(&Graph{})
}

/*
func (g *Graph) create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newEdge [3]int
	_ = json.NewDecoder(r.Body).Decode(&Graph{})
	for _, v := range newEdge {
		v = rand.Intn(10) // Mock ID - not safe
	}
	g.communications() = append(g.communications(), newEdge)
	json.NewEncoder(w).Encode(book)
}
*/
/*
func (g *Graph) update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range g.communications {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}
*/
