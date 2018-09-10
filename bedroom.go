package main

import "github.com/fatih/color"

type bedroom struct {
	room
}

func newBedroom() *bedroom {
	return &bedroom{room: room{
		short: "You are in a sunny, well arranged bedroom. The A/C is blowing cold air and a fan is spinning. " +
			"Tim is nowhere to be found. To the east is the living room.",
		desc: "The room has a bed, a night stand, a dresser with a clothes hamper next to it, " +
			"and some pretty illustrations on the walls. The door is open.",
		objects: map[string]*object{},
	}}
}

func (r *bedroom) data() *room {
	return &r.room
}

func (r *bedroom) onEnter() {
	if r.state == "" {
		color.HiYellow("You see a shadow move out of the corner of your eye. " +
			"Maybe a trick of the light? You don't see it again.")
		r.state = "visited"
	} else {
	}
}

func (r *bedroom) onExit() {}
func (r *bedroom) onLook() {}
