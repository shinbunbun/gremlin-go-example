package create

import (
	"fmt"
	"log"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

func CreateVertexes(g *gremlingo.GraphTraversalSource) ([]interface{}, error) {
	var resultVertexes []interface{}

	vertexes := sampleVertexes()
	for _, v := range vertexes {
		result, err := createVertex(g, v)
		if err != nil {
			return []interface{}{}, fmt.Errorf("create vertex error: %s", err.Error())
		}
		log.Printf("%#v", result.GetInterface())
		resultVertexes = append(resultVertexes, result.GetInterface())
	}

	return resultVertexes, nil
}

type vertex struct {
	label string
	name  string
}

func sampleVertexes() []vertex {
	return []vertex{
		{label: "user", name: "ユーザー1"},
		{label: "user", name: "ユーザー2"},
		{label: "user", name: "ユーザー3"},
		{label: "user", name: "ユーザー4"},
		{label: "user", name: "ユーザー5"},
		{label: "user", name: "ユーザー6"},
		{label: "user", name: "ユーザー7"},
		{label: "user", name: "ユーザー8"},
		{label: "user", name: "ユーザー9"},
		{label: "user", name: "ユーザー10"},
		{label: "post", name: "投稿1"},
		{label: "post", name: "投稿2"},
		{label: "post", name: "投稿3"},
	}
}

func createVertex(g *gremlingo.GraphTraversalSource, v vertex) (*gremlingo.Result, error) {
	return g.AddV(v.label).Property("name", v.name).Next()
}
