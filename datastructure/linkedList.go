package datastructure

import (
	"fmt"
	"container/list"
)

func TraverseLL(){
	myList := list.New()	//define new List
	myList.PushBack(1)
	myList.PushFront(2)	//2->1

	//operations on a node/Element
	element := myList.Front()
	myList.InsertAfter(5, element)
	myList.Remove(element)	//takes in a pointer to element

	//filter out unwanted element
	for element := myList.Front(); element != nil; element = element.Next() {
			if element.Value != 1 {
					myList.Remove(element)
			}
	}

	//iteration method 1: check until next node is nil
	for element := myList.Front(); element!=nil; element=element.Next(){
		fmt.Println(element.Value)
	}
}