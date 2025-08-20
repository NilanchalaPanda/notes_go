package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

// INTERFACES DEFINATION
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// EMBEDDED INTERFACES
type outputtable interface {
	saver
	// displayer
	Display()
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	// Directly using the NOTES package, instead of creating a Function for the same.
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}

	outputData(userNote)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return err
	}

	fmt.Println("Note saved successfully!")
	return nil
}

// ----------------- NOTE PACKAGE
// func getTodoData() string {
// 	return getUserInput("Todo text: ")
// }

// ----------------- NOTE PACKAGE
func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// Scope of improvement in this code will come in the form of INTERFACES!
