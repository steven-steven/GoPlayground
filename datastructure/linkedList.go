package datastructure

import (
	"container/list"
)

func TraverseLL(){
	myList := list.New()	//define new List
	myList.PushBack(1)
	myList.PushFront(2)	//2->1
}