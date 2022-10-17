//go:build !package
// +build !package

package create

import (
	"fmt"
	"log"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

func CreateEdges(g *gremlingo.GraphTraversalSource, vertexes []interface{}) error {
	edges := sampleEdges()
	for _, v := range edges {
		result, err := createEdge(g, vertexes[v.from-1], vertexes[v.to-1], v.label)
		if err != nil {
			return fmt.Errorf("create edges error: %s", err.Error())
		} else {
			log.Println(result.GetInterface())
		}
	}
	return nil
}

type edge struct {
	from  int
	to    int
	label string
}

func sampleEdges() []edge {
	return []edge{
		{from: 1, to: 2, label: "follow"},
		{from: 1, to: 3, label: "follow"},
		{from: 3, to: 4, label: "follow"},
		{from: 4, to: 5, label: "follow"},
		{from: 4, to: 6, label: "follow"},
		{from: 6, to: 7, label: "follow"},
		{from: 7, to: 8, label: "follow"},
		{from: 7, to: 9, label: "follow"},
		{from: 7, to: 10, label: "follow"},
		{from: 3, to: 11, label: "post"},
		{from: 10, to: 12, label: "post"},
		{from: 9, to: 13, label: "post"},
	}
}

func createEdge(g *gremlingo.GraphTraversalSource, vertex1 interface{}, vertex2 interface{}, label string) (*gremlingo.Result, error) {
	return g.AddE(label).From(vertex1).To(vertex2).Next()
}
