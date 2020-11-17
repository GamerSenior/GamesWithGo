package main

type choices struct {
	command     string
	description string
	node        *storyNode
	next        *choices
}

type storyNode struct {
	text string
}

func main() {

}
