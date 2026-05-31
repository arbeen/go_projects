package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Next steps:
/* type todo struct{
	ID int
	Text string
	Done bool
} */
const fileName = "todos.txt"

func addTodo(item string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(item + "\n")
	if err != nil {
		return err
	}
	return nil
}

func readTodo() ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if os.IsNotExist(err) {
		return []byte{}, nil
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func removeTodo(index int) error {
	found := false
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	tempFile, err := os.CreateTemp(".", "tmp-*")
	if err != nil {
		return err
	}
	tempName := tempFile.Name()
	defer tempFile.Close()
	defer os.Remove(tempName)

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tempFile)

	currentIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if currentIndex != index {
			if _, err := writer.WriteString(line + "\n"); err != nil {
				return err
			}
		}
		if currentIndex == index {
			found = true
		}
		currentIndex++
	}
	if !found {
		return fmt.Errorf("todo %d not found", index)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	if err := tempFile.Close(); err != nil {
		return err
	}
	if err := os.Remove(fileName); err != nil && !os.IsNotExist(err) {
		return err
	}
	return os.Rename(tempName, fileName)
}

func main() {
	var action string

	fmt.Println("Actions:")
	fmt.Println("1. Add new todo")
	fmt.Println("2. Remove a todo")
	fmt.Println("3. Print todos")

	fmt.Scanln(&action)

	switch action {
	case "1":
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter new todo")
		if scanner.Scan() {
			if err := addTodo(scanner.Text()); err != nil {
				log.Fatal(err)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	case "2":
		var indexInput string
		todos, err := readTodo()
		if err != nil {
			log.Fatal(err)
		}
		if len(todos) == 0 {
			fmt.Println("No todos found")
			return
		}
		splitTodo := bytes.Split(todos, []byte("\n"))
		for i, todo := range splitTodo {
			fmt.Printf("%v. %v\n", i, string(todo))
		}
		// fmt.Println(string(todos))
		fmt.Println("Enter todo # to remove")
		fmt.Scanln(&indexInput)
		// Atoi uses default decimal base, where ParseInt takes (s string, base int, bitSize int) (int64, error)
		removeIndex, err := strconv.Atoi(indexInput)
		if err != nil {
			log.Fatal(err)
		}
		err = removeTodo(removeIndex)
		if err != nil {
			log.Fatal(err)
		}
	case "3":
		todos, err := readTodo()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(todos))
	default:
		fmt.Println("Bye.")
	}
}
