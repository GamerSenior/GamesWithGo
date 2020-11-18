package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choices struct {
	command     string
	description string
	nextNode    *storyNode
	nextChoice  *choices
}

type storyNode struct {
	text    string
	choices *choices
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choices{cmd, description, nextNode, nil}

	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.command, ": ", currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

func (node *storyNode) executeCommand(cmd string) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.EqualFold(currentChoice.command, cmd) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("Sorry, I didn't understand that")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCommand(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{
		text: `You are lost in a gigantic building, filled with computers.
		Nobody can be seen around you.`,
	}

	darkRoom := storyNode{
		text: `You find yourself in a dark room.`,
	}

	darkRoomLit := storyNode{
		text: `The dark room is now lit and you can see a door at the end of the room`,
	}

	monster := storyNode{
		text: `While walking in the dark, a monster eats you`,
	}

	trap := storyNode{
		text: `While walking through the room, you fall into a trap and the spikes kill you`,
	}

	treasure := storyNode{
		text: `You found a chest, full of treasures`,
	}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &monster)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("B", "Go back", &monster)
	darkRoom.addChoice("L", "Turn on the lights", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	start.play()

	fmt.Println("The End.")
}
