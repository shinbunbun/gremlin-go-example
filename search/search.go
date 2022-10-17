package search

import (
	"fmt"
	"log"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

// グラフに登録されているユーザーの名前一覧を取得
func GetAllUsersName(g *gremlingo.GraphTraversalSource) {
	res, err := g.V().HasLabel("user").Values("name").ToList()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(res)
	}
}

// ユーザー1に紐づくuserの名前一覧を取得
func GetRelatedUsersName(g *gremlingo.GraphTraversalSource, userName string) {
	res, err := g.V().Has("user", "name", userName).Out().HasLabel("user").Values("name").ToList()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(res)
	}
}

// ユーザー7がフォローしているユーザーの投稿を取得
func GetFollowUsersPosts(g *gremlingo.GraphTraversalSource, userName string) {
	res, err := g.V().Has("user", "name", userName).Out().HasLabel("user").Out().HasLabel("post").Values("name").ToList()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(res)
	}
}

// ユーザー3から枝方向に再帰的に探索し、ユーザーの名前一覧を取得
func GetUsersRecursive(g *gremlingo.GraphTraversalSource, userName string) {
	res, err := g.V().Has("user", "name", userName).Repeat(g.GetGraphTraversal().Out()).Emit(g.GetGraphTraversal().HasLabel("user")).Values("name").ToList()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(res)
	}
}

// 再帰で2階層下（ユーザー3がフォローしているユーザーがフォローしているユーザー）まで取得
func GetUsersRecursive2Step(g *gremlingo.GraphTraversalSource, userName string) {
	res, err := g.V().Has("user", "name", "ユーザー3").Repeat(g.GetGraphTraversal().Out()).Times(2).Emit(g.GetGraphTraversal().HasLabel("user")).Values("name").ToList()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(res)
	}
}
