package main

import "github.com/fatih/color"

/*

Ideas:

Look around house
Find cryptic note "The monster likes candy and darkness. Bring a swimsuit."?
Secret tunnel in office
Gather tools (not in order!):
* flashlight
* hammer
* water bottle
* candy
* swimsuit
* cool leather jacket
Use flashlight in tunnel
Maze-like...
Tunnel creatures
Tunnel leads to another dimension - tiny planet
Need to spread out locations
Useless but funny point system
Silly porn tape

*/

var story = map[string]iroom{
	"bedroom":    newBedroom(),
	"livingroom": newLivingroom(),
	"kitchen":    newKitchen(),
	"office":     newOffice(),
}

var objects = map[string]*object{
	"swimsuit": &object{
		roomDesc: "One of the dresser drawers is slightly pulled out.",
		pickupDesc: "You grab your swimsuit, a fun " + color.HiBlackString("black") + " & " +
			color.HiWhiteString("white") + " two-piece. Maybe it will come in handy.",
		lookDesc: "You pull the drawer out further and notice that Tim's swim trunks are missing. " +
			"This is odd. You stare at your own swimsuit while trying to figure out what to do.",
		invDesc:  "A crumpled up two-piece swimsuit is stuffed in your back pocket. Maybe it will come in handy.",
		synonyms: []string{"swim suit", "bikini", "clothes", "clothing", "trunks", "dresser", "drawer", "dresser drawer"},
		gettable: true,
	},
	"jacket": &object{
		roomDesc: "Some jackets are hanging on a tasteful wooden hook bar by the door.",
		pickupDesc: "You notice your super cool, well worn " + color.HiBlackString("black leather jacket") + " and put it on. " +
			"It envelops your body like a soft glove and makes you feel like cracking your knuckles " +
			"and making a witty quip.",
		lookDesc: "One of your " + color.HiBlackString("black jackets") + " seems particularly enticing.",
		invDesc:  "You are wearing a really cool " + color.HiBlackString("leather jacket."),
		synonyms: []string{"jackets", "hooks", "hook", "hook bar", "wooden hook"},
		gettable: true,
	},
	"drawer": &object{
		roomDesc: "You remember there are some emergency supplies in a drawer by the pantry.",
		lookDesc: "You think the supplies drawer may have some lightbulbs, a screwdriver, and a flashlight.",
		synonyms: []string{"supplies", "emergency drawer", "supply drawer", "emergency supplies drawer"},
		usable:   true,
	},
	"flashlight": &object{
		roomDesc: "There is a small " + color.HiGreenString("green LED flashlight") + " in a drawer by the pantry.",
		lookDesc: "You look at the front of the flashlight and turn it on. The light is blinding! " +
			"You turn it off quickly and blink a bunch. :(",
		pickupDesc: "You put the flashlight in your pocket. Maybe it will come in handy later. Or blind you.",
		invDesc:    "You have a really bright LED flashlight in your pocket.",
		synonyms:   []string{"led flashlight", "green led flashlight"},
		gettable:   true,
	},
	"candy": &object{
		roomDesc: "There may be some " + color.HiGreenString("candy") + " in the pantry.",
		lookDesc: "You rummage in the pantry and find a box of Annie's Tropical Fruit Snack " + color.HiGreenString("candy") + ". " +
			"A faint sugary smells wafts into your nostrils. Mmmmmmmm",
		pickupDesc: "One of these fruit snack " + color.HiGreenString("candy") + " packets will fit " +
			"perfectly in your pocket. Let's hope they don't melt!",
		invDesc: "A packet of Annie's Fruit Snack " + color.HiGreenString("candy") + " is slowly melting into a sweet " +
			"blob in your pocket.",
		synonyms: []string{"snack", "snacks", "fruit snacks", "annie's fruit snack candy"},
		gettable: true,
	},
}

func init() {
	story["bedroom"].data().exitEast = story["livingroom"]
	story["bedroom"].data().objects["swimsuit"] = objects["swimsuit"]
	story["livingroom"].data().exitWest = story["bedroom"]
	story["livingroom"].data().exitSouth = story["kitchen"]
	story["livingroom"].data().exitEast = story["office"]
	story["livingroom"].data().objects["jacket"] = objects["jacket"]
	story["kitchen"].data().exitNorth = story["livingroom"]
	story["kitchen"].data().objects["drawer"] = objects["drawer"]
	story["kitchen"].data().objects["candy"] = objects["candy"]
	story["office"].data().exitWest = story["livingroom"]
	objects["drawer"].action = openDrawer

}

func openDrawer(obj *object, verb string) {
	if obj.state == "checked" {
		writeLn(color.YellowString("You've already checked the drawer."))
		return
	}
	writeLn("You open the drawer and notice a small " + color.HiGreenString("green LED flashlight") +
		" sitting right on top.")
	obj.state = "checked"
	obj.picked = true
	story["kitchen"].data().objects["flashlight"] = objects["flashlight"]
}
