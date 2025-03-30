package main

import (
	"ds/hashtable"
	"ds/queue"
	"ds/stack"
	"fmt"
)

func main() {
	miStack := stack.NewStack()
	miQueue := queue.NewQueue()
	miHashtable := hashtable.NewHashTable(10)
	fmt.Println("---------------Pruebas de stack---------------")
	for i := 0; i < 4; i++ {

		miStack.Push(i)
		fmt.Println("Stack Elemento agregado:", i)
		miStack.ShowStack()
	}

	for !miStack.IsEmpty() {
		valorStack, err := miStack.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Stack Elemento eliminado:", valorStack)
		}
		miStack.ShowStack()
	}

	fmt.Println("---------------Pruebas de queue---------------")
	for i := 0; i < 4; i++ {

		miQueue.Push(i)
		fmt.Println("Queue Elemento agregado:", i)
		miQueue.ShowQueue()
	}
	for !miQueue.IsEmpty() {
		valorQueue, err := miQueue.Dequeue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Queue Elemento eliminado:", valorQueue)
		}
		miQueue.ShowQueue()
	}

	fmt.Println("---------------Pruebas de hashtable---------------")

	miHashtable.Set("nombre", "Rafha")
	miHashtable.Set("novia", "Ale")

	miHashtable.Print()
	err := miHashtable.Delete("nombre")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hastable Elemento eliminado: nombre")
	}

	miHashtable.Print()

}
