package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

//URL - API URL
const URL = "https://25.ms/posts/"

//CommentJSON - API JSON format
type CommentJSON struct {
	UserID int
	ID     int
	Title  string
	Body   string
}

func insertCommentFromAPI(node *TreeNode) {
	defer wg.Done()
	res, err := http.Get(URL + strconv.Itoa(node.id))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data CommentJSON
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	node.body = data.Title
}

//TreeNode - structure of tree node
type TreeNode struct {
	id      int
	body    string
	replies []*TreeNode
}

func treeTreversal(node *TreeNode) {
	wg.Add(1)
	go insertCommentFromAPI(node)

	if node.replies == nil {
		return
	}
	for _, val := range node.replies {
		treeTreversal(val)
	}
}

func traverse(node *TreeNode) {
	fmt.Println(node.id, node.body)

	if node.replies == nil {
		return
	}
	for _, val := range node.replies {
		traverse(val)
	}
}

var wg sync.WaitGroup

func main() {
	t0 := time.Now()

	var tree = TreeNode{id: 1, body: "111", replies: []*TreeNode{
		{id: 2, body: "222", replies: []*TreeNode{
			{id: 3, body: "333", replies: nil},
			{id: 4, body: "444", replies: nil},
		}},
		{id: 5, body: "555", replies: nil},
	}}
	treeTreversal(&tree)

	wg.Wait()
	traverse(&tree)

	t1 := time.Now()
	fmt.Printf("Elapsed time: %v", t1.Sub(t0))
}
