package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Todo struct {
	ID   int
	Text string
	Done bool
}

func addTodo(todo string) error {
	return nil
}

// TODO: mark as done
// TODO: printTodos
// TODO: deleteTodo

func main() {
	fmt.Println("Choose an action:")
	fmt.Println("1. Add a todo")
	fmt.Println("2. Mark todo done")
	fmt.Println("3. List all todos")
	fmt.Println("4. Delete a todo")
	fmt.Println("*. Exit")
	var action string
	fmt.Scan(&action)

	switch action {
	case "1":
		//add a todo
		fmt.Println("Enter todo: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			if err := addTodo(scanner.Text()); err != nil {
				log.Fatal(err)
			}
		}
	}
}
