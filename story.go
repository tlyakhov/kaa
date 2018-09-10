package main

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
}

var objects = map[string]*object{
	"jacket": &object{
		roomDesc: "Some jackets are hanging on a tasteful wooden hook bar by the door.",
		pickupDesc: "You notice your super cool, well worn black leather jacket and put it on. " +
			"It envelops your body like a soft glove and makes you feel like cracking your knuckles " +
			"and making a witty quip.",
		lookDesc: "One of your black jackets seems particularly enticing.",
		invDesc:  "You are wearing a really cool leather jacket.",
		synonyms: []string{"jackets", "hooks", "hook", "hook bar", "wooden hook"},
	},
}

func init() {
	story["bedroom"].data().exitEast = story["livingroom"]
	story["livingroom"].data().exitWest = story["bedroom"]
	story["livingroom"].data().objects["jacket"] = objects["jacket"]

}
