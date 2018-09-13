package main

import "github.com/fatih/color"

type livingroom struct {
	room
}

func newLivingroom() *livingroom {
	return &livingroom{room: room{
		short: "You are in a large living room with a fancy home theater system and a black leather couch. " +
			"To the west is the bedroom. To the south is the kitchen.",
		desc: "The room is large, but quite cozy. An unused fireplace has a giant log crammed into it. " +
			"Many colorful books line a row of shelves behind the couch. An obscenely high-end turntable sits on top of them. " +
			"You feel a little afraid of it.",
		objects: map[string]*object{},
	}}
}

func (r *livingroom) data() *room {
	return &r.room
}

func (r *livingroom) onEnter() {
	if r.state == "" {
		writeLn(color.HiBlackString("Tim's sunglasses and flip-flops are missing. " +
			"Maybe he went out to get something?"))
		r.state = "visited"
	} else {
	}
}

func (r *livingroom) onExit() {}
func (r *livingroom) onLook() {}
