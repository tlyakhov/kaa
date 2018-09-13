package main

import "github.com/fatih/color"

type office struct {
	room
}

func newOffice() *office {
	return &office{room: room{
		short: "You are in Tim's office/studio. It's quiet other than the soft humming of the server. " +
			"It's a little too messy for your liking, but you've learned to live with it. " +
			"To the west is the living room.",
		desc: "The office is full of music gear: a full drumkit, many guitars, and acoustic " +
			"panels covering the walls. There's also a custom wooden desk with a giant monitor on it. " +
			"You try not to touch anything.",
		objects: map[string]*object{},
	}}
}

func (r *office) data() *room {
	return &r.room
}

func (r *office) onEnter() {
	if r.state == "" {
		writeLn(color.HiBlackString("A strange damp smell is wafting from the direction " +
			"of the office bathroom. It's a bit alarming, but you're not sure if you want to " +
			"investigate."))
		r.state = "visited"
	} else {
	}
}

func (r *office) onExit() {}
func (r *office) onLook() {}
