package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

const MaxLineLength = 70

type object struct {
	state      string
	roomDesc   string
	lookDesc   string
	invDesc    string
	pickupDesc string
	synonyms   []string
	action     func(obj *object, verb string)
	picked     bool
	usable     bool
	gettable   bool
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

func listObjectRoomDesc(room iroom) string {
	desc := ""
	for _, obj := range room.data().objects {
		if obj.picked {
			continue
		}
		desc += " " + obj.roomDesc
	}
	return desc
}
func enterRoom(room iroom) {
	currentRoom = room
	writeLn(room.data().short + listObjectRoomDesc(room))
	room.onEnter()
}

func actionLookAround() {
	writeLn(currentRoom.data().desc + listObjectRoomDesc(currentRoom))
	currentRoom.onLook()
}

func write(str string, a ...interface{}) {
	result := fmt.Sprintf(str, a...)
	count := 0
	lastSpace := -1
	for i, c := range result {
		if count > MaxLineLength && lastSpace != -1 {
			result = result[:lastSpace] + "\n" + result[lastSpace+1:]
			count = 0
			lastSpace = -1
		}
		if c == ' ' {
			lastSpace = i
		} else if c == '\n' {
			lastSpace = -1
			count = -1
		}
		if c == '\x1b' {
			count -= 4
		}
		count++
	}
	fmt.Print(result)
}

func writeLn(str string, a ...interface{}) {
	write(str+"\n", a...)
}

func actionQuit() {
	writeLn(color.HiWhiteString("Sorry to see you go!"))
	os.Exit(0)
}

func actionHelp() {
	writeLn(color.HiWhiteString("Welcome to Kayla's Anniversary adventure."))
	writeLn(color.WhiteString("(c)1986 Verdugo Hills Wetware"))
	writeLn("")
	writeLn(color.HiCyanString("Your name is Kayla, and you've got a mystery to solve!"))
	writeLn(color.WhiteString("Describe what you want to do in the form of a command (e.g. start with a verb)."))
	writeLn(color.WhiteString("No punctuation, capitalization, or complex phrasing. Some examples of what you can say:"))
	writeLn(color.HiCyanString("help") + color.HiBlackString(" - show this text again."))
	writeLn(color.HiCyanString("quit") + color.HiBlackString(" - stop playing. 'exit' also works."))
	writeLn(color.HiCyanString("inventory") + color.HiBlackString(" - list your inventory."))
	writeLn(color.HiCyanString("go [north|south|east|west]") + color.HiBlackString(" - move to another place. You can also just type the direction without 'go'."))
	writeLn(color.HiCyanString("look/look around") + color.HiBlackString(" - take a more detailed look at the place you're in."))
	writeLn(color.HiCyanString("pick up [object]") + color.HiBlackString(" - take something. 'grab', 'get', and other variations also work."))
	writeLn(color.HiCyanString("look at [object]") + color.HiBlackString(" - take a more detailed look at an object."))
	writeLn(color.HiCyanString("use [object]") + color.HiBlackString(" - try to manipulate something."))
	writeLn("")
	writeLn("")
	writeLn("")
}

func actionInventory() {
	if len(inventory) == 0 {
		writeLn("You don't have anything in your inventory.")
		return
	}
	inv := ""
	for _, obj := range inventory {
		inv += obj.invDesc + " "
	}
	writeLn(inv + "\n")
}

func matchObject(target string) (string, *object) {
	for name, obj := range inventory {
		if target == name {
			return name, obj
		}
		for _, syn := range obj.synonyms {
			if target != syn {
				continue
			}
			return name, obj
		}
	}
	for name, obj := range currentRoom.data().objects {
		if target == name {
			return name, obj
		}
		for _, syn := range obj.synonyms {
			if target != syn {
				continue
			}
			return name, obj
		}
	}
	return "", nil
}
func actionLookAt(target string) {
	_, obj := matchObject(target)
	if obj == nil {
		writeLn(color.RedString("Nothing special about the " + target + " here. :("))
		return
	} else if obj.picked {
		writeLn(obj.invDesc)
		return
	}
	writeLn(obj.lookDesc)
}

func actionGo(target string) {
	switch target {
	case "north":
		if currentRoom.data().exitNorth == nil {
			writeLn(color.RedString("There's nothing to the north here!"))
			return
		}
		currentRoom.onExit()
		enterRoom(currentRoom.data().exitNorth)
	case "east":
		if currentRoom.data().exitEast == nil {
			writeLn(color.RedString("There's nothing to the east here!"))
			return
		}
		currentRoom.onExit()
		enterRoom(currentRoom.data().exitEast)
	case "west":
		if currentRoom.data().exitWest == nil {
			writeLn(color.RedString("There's nothing to the west here!"))
			return
		}
		currentRoom.onExit()
		enterRoom(currentRoom.data().exitWest)
	case "south":
		if currentRoom.data().exitSouth == nil {
			writeLn(color.RedString("There's nothing to the south here!"))
			return
		}
		currentRoom.onExit()
		enterRoom(currentRoom.data().exitSouth)
	default:
		writeLn(color.RedString("I don't know what you're trying to do! :("))
	}
}

func actionGet(target string) {
	name, obj := matchObject(target)
	if obj == nil {
		writeLn(color.RedString("No reason to pick up the " + target + " here. :("))
		return
	} else if obj.picked {
		writeLn(color.YellowString("Already picked it up!"))
		return
	} else if !obj.gettable {
		writeLn(color.YellowString("Can't pick up the " + target + ". :("))
		return
	}
	writeLn(obj.pickupDesc)
	obj.picked = true
	inventory[name] = obj
	if obj.action != nil {
		obj.action(obj, "get")
	}
}

func actionUse(target string) {
	_, obj := matchObject(target)
	if obj == nil {
		writeLn(color.RedString("Don't know how to use the " + target + " here. :("))
		return
	} else if !obj.usable {
		writeLn(color.RedString("Don't know how to use the " + target + ". :("))
		return
	}
	obj.action(obj, "use")
}

func userInput() {
	generalError := "I don't know what you're trying to do!"
	scanner := bufio.NewScanner(os.Stdin)
	writeLn("")
	fmt.Print(color.HiBlackString("What next? "))
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
			writeLn(color.RedString(generalError))
			return
		}
		actionGo(split[1])
	case "north", "south", "east", "west":
		actionGo(split[0])
	case "pick":
		if len(split) > 2 && split[1] == "up" {
			actionGet(strings.Join(split[2:], " "))
		} else {
			writeLn(color.RedString("Not sure what you mean... Did you want to pick something up?"))
		}
	case "get", "grab", "take":
		actionGet(strings.Join(split[1:], " "))
	case "use":
		actionUse(strings.Join(split[1:], " "))
	default:
		color.RedString(generalError)
	}
}

func main() {
	actionHelp()
	enterRoom(story["bedroom"])
	for {
		userInput()
	}
}
