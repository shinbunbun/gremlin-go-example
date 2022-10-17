package main

import (
	"fmt"
	"gremlin-go-example/create"
	"gremlin-go-example/search"
	"log"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

func main() {
	host := "ws://graph:8182/gremlin"

	driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection(host)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer driverRemoteConnection.Close()
	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	// vertex作成
	vertexes, err := create.CreateVertexes(g)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println()

	// edge作成
	err = create.CreateEdges(g, vertexes)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("\nグラフに登録されているユーザーの名前一覧を取得")
	search.GetAllUsersName(g)

	fmt.Println("\nユーザー1に紐づくuserの名前一覧を取得")
	search.GetRelatedUsersName(g, "ユーザー1")

	fmt.Println("\nユーザー7がフォローしているユーザーの投稿を取得")
	search.GetFollowUsersPosts(g, "ユーザー7")

	fmt.Println("\nユーザー3から枝方向に再帰的に探索し、ユーザーの名前一覧を取得")
	search.GetUsersRecursive(g, "ユーザー3")

	fmt.Println("\n再帰で2階層下（ユーザー3がフォローしているユーザーがフォローしているユーザー）まで取得")
	search.GetUsersRecursive2Step(g, "ユーザー3")

	fmt.Println()
}
