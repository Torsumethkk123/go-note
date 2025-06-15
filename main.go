package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Note struct {
	Content string
}

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// show help options
func showHelp() {
	fmt.Println("help - show all avaliable options")
	fmt.Println("add - add a note")
	fmt.Println("edit - edit the note")
	fmt.Println("remove - remove the note")
	fmt.Println("exit - exit the program")
}

func main() {
	// setup
	reader := bufio.NewReader(os.Stdin)
	notes := []Note{}

	for true {
		command := ""

		// show realtime exist note
		fmt.Println("------------------------")

		fmt.Println("Go - Notes")

		fmt.Println("------------------------")
		
		if len(notes) == 0 {
			fmt.Println("> No notes")
		} else {
			for i, note := range notes {
				fmt.Println("> " + strconv.Itoa(i + 1) + ") " + note.Content)
			}
		}

		fmt.Println("------------------------")

		// check user action e.g. add edit or remove
		fmt.Println("Enter a option (Use help to see avaliable options or use exit for exit the program ): ")
		fmt.Scanln(&command)

		if command == "help" {
			showHelp()
			helpCommand := ""
			fmt.Println("Enter e to exit help menu")
			for helpCommand != "e" {
				fmt.Scanln(&helpCommand)
			}
		} else if command == "add" {
			fmt.Println("Enter note content here: ")
			content, _ := reader.ReadString('\n')
			content = strings.TrimSpace(content)
			notes = append(notes, Note{Content: content})
		} else if command == "edit" {
			if len(notes) != 0 {
				newNotes := []Note{}
				editId := 0
				fmt.Println("Enter id of note that you want to edit: ")
				fmt.Scanln(&editId)
				fmt.Println("Enter new content: ")
				newContent, _ := reader.ReadString('\n') 
				newContent = strings.TrimSpace(newContent)
				if editId != 0 {
					for i, note := range notes {
						if i + 1 != editId {
							newNotes = append(newNotes, note)
						} else {
							newNotes = append(newNotes, Note{ Content: newContent })
						}
					}
					notes = newNotes
				}
			}
		} else if command == "remove" {
			if len(notes) != 0 {
				newNotes := []Note{}
				removeId := 0
				fmt.Println("Enter id of note that you want to remove: ")
				fmt.Scanln(&removeId)
				if removeId != 0 {
					for i, note := range notes {
						if i + 1 != removeId {
							newNotes = append(newNotes, note)
						}
					}
					notes = newNotes
				}
			} else {
				fmt.Println("Can't remove no notes here!")
			}
		} else if command == "exit" {
			fmt.Println("Program is ended.")
			clearTerminal()
			break
		} else {
			continue
		}

		clearTerminal()
	}
}