package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type object struct {
	roomDesc   string
	lookDesc   string
	invDesc    string
	pickupDesc string
	synonyms   []string
	action     func(verb string)
}

type room struct {
	short     string
	desc      string
	state     string
	exitNorth iroom
	exitSouth iroom
	exitWest  iroom
	exitEast  iroom
	objects   map[string]*object
}

type iroom interface {
	onEnter()
	onLook()
	onExit()
	data() *room
}

var currentRoom iroom
var inventory = map[string]*object{}

func enterRoom(room iroom) {
	currentRoom = room
	fmt.Println(room.data().short)
	for _, obj := range room.data().objects {
		fmt.Println(obj.roomDesc)
	}
	room.onEnter()
}

func actionLookAround() {
	fmt.Println(currentRoom.data().desc)
	for _, obj := range currentRoom.data().objects {
		fmt.Println(obj.roomDesc)
	}
	currentRoom.onLook()
}

func actionQuit() {
	color.HiWhite("Sorry to see you go!")
	os.Exit(0)
}

func actionHelp() {
	color.HiWhite("Welcome to Kayla's Anniversary adventure.")
	color.White("(c)1986 Verdugo Hills Wetware")
	color.White("")
	color.HiGreen("Your name is Kayla, and you've got a mystery to solve!")
	color.White("Describe what you want to do in the form of a command (e.g. start with a verb).")
	color.White("No punctuation, capitalization, or complex phrasing. Some examples of what you can say:")
	fmt.Println(color.HiWhiteString("help") + color.WhiteString(" - show this text again."))
	fmt.Println(color.HiWhiteString("quit") + color.WhiteString(" - stop playing. 'exit' also works."))
	fmt.Println(color.HiWhiteString("inventory") + color.WhiteString(" - list your inventory."))
	fmt.Println(color.HiWhiteString("go [north|south|east|west]") + color.WhiteString(" - move to another place. You can also just type the direction without 'go'."))
	fmt.Println(color.HiWhiteString("pick up [object]") + color.WhiteString(" - take something. 'grab', 'get', and other variations also work."))
	fmt.Println(color.HiWhiteString("look/look around") + color.WhiteString(" - take a more detailed look at the place you're in."))
	color.White("")
	color.White("")
	color.White("")
}

func actionInventory() {
	if len(inventory) == 0 {
		fmt.Println("You don't have anything in your inventory.")
		return
	}
	for _, obj := range inventory {
		fmt.Print(obj.invDesc)
	}
	fmt.Println()
}

func matchObject(target string) *object {
	for name, obj := range currentRoom.data().objects {
		if target == name {
			return obj
		}
		for _, syn := range obj.synonyms {
			if target != syn {
				continue
			}
			return obj
		}
	}
	return nil
}
func actionLookAt(target string) {
	obj := matchObject(target)
	if obj == nil {
		fmt.Println(color.RedString("Don't know what " + target + " is. :("))
		return
	}
	fmt.Println(obj.lookDesc)
}

func actionGo(target string) {
	switch target {
	case "north":
		if currentRoom.data().exitNorth == nil {
			color.Red("There's nothing to the north here!")
			return
		}
		currentRoom.onExit()
		enterRoom(currentRoom.data().exitNorth)
	case "east":
		if currentRoom.data().exitEast == nil {
			color.Red("There's nothing to the east here!")
			return
		}
		currentRoom.onExit()
		enterRoom(currentRoom.data().exitEast)
	case "west":
		if currentRoom.data().exitWest == nil {
			color.Red("There's nothing to the west here!")
			return
		}
		currentRoom.onExit()
		enterRoom(currentRoom.data().exitWest)
	case "south":
		if currentRoom.data().exitSouth == nil {
			color.Red("There's nothing to the south here!")
			return
		}
		currentRoom.onExit()
		enterRoom(currentRoom.data().exitSouth)
	default:
		color.Red("I don't know what you're trying to do! :(")
	}
}

func userInput() {
	generalError := "I don't know what you're trying to do!"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println()
	fmt.Print(color.WhiteString("What next? "))
	read := scanner.Scan()
	if !read {
		os.Exit(0)
	}
	input := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if len(input) == 0 {
		return
	}
	split := strings.Split(input, " ")
	switch split[0] {
	case "quit", "exit":
		actionQuit()
	case "help", "?":
		actionHelp()
	case "inventory":
		actionInventory()
	case "look":
		if len(split) == 1 || split[1] == "around" {
			actionLookAround()
		} else if split[1] == "at" {
			actionLookAt(strings.Join(split[2:], " "))
		}
	case "go":
		if len(split) == 1 {
			color.Red(generalError)
			return
		}
		actionGo(split[1])
	case "north", "south", "east", "west":
		actionGo(split[0])
	default:
		color.Red(generalError)
	}
}

func main() {
	actionHelp()
	enterRoom(story["bedroom"])
	for {
		userInput()
	}
}
