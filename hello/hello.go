package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Arbin")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	people := []string {
		"arbin", 
		"tim", 
		"josh",
	}
	messages, err := greetings.HelloPeople(people)
	if err != nil{
		log.Fatal(err)
	}
	for _, message := range messages{
		fmt.Println(message)
	}

}
