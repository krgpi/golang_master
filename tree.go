package main

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	key   int
	value int
	left  *node
	right *node
}

type tree struct {
	root *node
}

//二分探索木を生成する
func maketree(parentNode *node, newKey int, newValue int) {
	var childNode node
	childNode.key = newKey
	childNode.value = newValue
	if newValue >= parentNode.value {
		if parentNode.right == nil {
			parentNode.right = &childNode
		} else {
			maketree(parentNode.right, newKey, newValue)
		}
	} else {
		if parentNode.left == nil {
			parentNode.left = &childNode
		} else {
			maketree(parentNode.left, newKey, newValue)
		}
	}
}

//findNumber 深さ優先探索
func findNumber(targetNode *node, targetNumber int) {
	if targetNode.value == targetNumber {
		fmt.Printf("target number[%v] found at key[%v]\n", targetNumber, targetNode.key)
	} else if targetNumber > targetNode.value {
		if targetNode.right != nil {
			findNumber(targetNode.right, targetNumber)
		} else {
			fmt.Printf("target number[%v] not found.\n", targetNumber)
		}
	} else {
		if targetNode.left != nil {
			findNumber(targetNode.left, targetNumber)
		} else {
			fmt.Printf("target number[%v] not found.\n", targetNumber)
		}
	}
}

func main() {
	//乱数生成器
	var valueArray = []int{500}
	rand.Seed(time.Now().UnixNano())
	for index := 0; index < 1000; index = index + 1 {
		valueArray = append(valueArray, rand.Intn(1000))
	}

	var myTree tree
	var firstNode node
	myTree.root = &firstNode

	//二分探索木の生成
	for index := 0; index < len(valueArray); index = index + 1 {
		maketree(myTree.root, index, valueArray[index])
	}

	//探索
	findNumber(myTree.root, 100)
}
